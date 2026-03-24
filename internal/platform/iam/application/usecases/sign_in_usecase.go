package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/ports"
)

type SignIn struct {
	operatorRepository ports.OperatorRepository
}

func NewSignIn(operatorRepository ports.OperatorRepository) *SignIn {
	return &SignIn{
		operatorRepository: operatorRepository,
	}
}

func (r *SignIn) Execute(ctx context.Context, cmd *commands.SignIn) (*aggregates.Operator, error) {
	operator, err := r.operatorRepository.FindOneByUsername(ctx, cmd.Username)

	if err != nil {
		return nil, err
	}

	if operator == nil {
		return nil, derrors.ErrOperatorNotFound
	}

	if !operator.VerifyPIN(cmd.PIN) {
		return nil, derrors.ErrOperatorIsNotMemberOfOrg
	}

	return operator, nil
}
