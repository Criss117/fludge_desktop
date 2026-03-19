package usecases

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
)

type SignUpUseCase struct {
	operatorRepo ports.OperatorRepository
	appStateRepo ports.AppStateRepository
}

func NewSignUpUseCase(
	operatorRepo ports.OperatorRepository,
	appStateRepo ports.AppStateRepository,
) *SignUpUseCase {
	return &SignUpUseCase{
		operatorRepo: operatorRepo,
		appStateRepo: appStateRepo,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, cmd *commands.SignUpCommand) (*aggregates.AppState, error) {
	existingOperator, errByUsername := uc.operatorRepo.FindOneByUsername(ctx, cmd.Username)

	if errByUsername != nil {
		return nil, errByUsername
	}

	if existingOperator != nil {
		return nil, derrors.ErrOperatorAlreadyExistsByUsername
	}

	existingOperator, errByEmail := uc.operatorRepo.FindOneByEmail(ctx, cmd.Email)

	if errByEmail != nil {
		return nil, errByEmail
	}

	if existingOperator != nil {
		return nil, derrors.ErrOperatorAlreadyExistsByEmail
	}

	newOperator, err := aggregates.NewOperator(
		cmd.Name,
		cmd.Username,
		cmd.Email,
		cmd.PIN,
		true,
		nil,
	)

	if err != nil {
		return nil, err
	}

	errCreating := uc.operatorRepo.Create(ctx, newOperator)

	if errCreating != nil {
		return nil, errCreating
	}

	currentAppState, errAppState := uc.appStateRepo.Get(ctx)

	if errAppState != nil {
		return nil, errAppState
	}

	currentAppState.SetActiveOperator(existingOperator)

	errUpdateAppState := uc.appStateRepo.Update(ctx, currentAppState)

	if errUpdateAppState != nil {
		return nil, errUpdateAppState
	}

	return currentAppState, nil
}
