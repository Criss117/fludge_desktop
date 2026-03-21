package usecases

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
)

type SignInUseCase struct {
	operatorRepo ports.OperatorRepository
	appStateRepo ports.AppStateRepository
	teamRepo     ports.TeamRepository
}

func NewSignInUseCase(
	operatorRepo ports.OperatorRepository,
	appStateRepo ports.AppStateRepository,
) *SignInUseCase {
	return &SignInUseCase{
		operatorRepo: operatorRepo,
		appStateRepo: appStateRepo,
	}
}

func (uc *SignInUseCase) Execute(ctx context.Context, cmd *commands.SignIn) (*aggregates.AppState, error) {
	existingOperator, errByUsername := uc.operatorRepo.FindOneByUsername(ctx, cmd.Username)

	if errByUsername != nil {
		return nil, errByUsername
	}

	if existingOperator == nil {
		return nil, derrors.ErrOperatorNotFound
	}

	if !existingOperator.ValidatePin(cmd.PIN) {
		return nil, derrors.ErrInvalidCredentials
	}

	currentAppState, errAppState := uc.appStateRepo.Get(ctx)

	if errAppState != nil {
		return nil, errAppState
	}

	currentAppState.SetActiveOperator(existingOperator)

	err := uc.appStateRepo.Update(ctx, currentAppState)

	if err != nil {
		return nil, err
	}

	return currentAppState, nil
}
