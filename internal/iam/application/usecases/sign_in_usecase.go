package usecases

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
)

type SignInResponse struct {
	Operator *aggregates.Operator
	Teams    []*aggregates.Team
}

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

func (uc *SignInUseCase) Execute(ctx context.Context, cmd *commands.SignInCommand) (*SignInResponse, error) {
	existingOperator, errByUsername := uc.operatorRepo.FindOneByUsername(ctx, cmd.Username)

	if errByUsername != nil {
		return nil, errByUsername
	}

	if existingOperator != nil {
		return nil, derrors.ErrOperatorNotFound
	}

	if !existingOperator.ValidatePin(cmd.PIN) {
		return nil, derrors.ErrInvalidCredentials
	}

	if existingOperator.IsRoot() {
		return &SignInResponse{Operator: existingOperator}, nil
	}

	teams, errTeams := uc.teamRepo.FindAllByOperatorId(ctx, existingOperator.ID())

	if errTeams != nil {
		return nil, errTeams
	}

	return &SignInResponse{
		Operator: existingOperator,
		Teams:    teams,
	}, nil
}
