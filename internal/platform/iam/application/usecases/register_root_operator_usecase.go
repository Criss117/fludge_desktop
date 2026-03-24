package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/platform/iam/domain/valueobjects"
	"desktop/internal/shared/db/dbutils"
)

type RegisterRootOperator struct {
	operatorRepository ports.OperatorRepository
}

func NewRegisterRootOperator(operatorRepository ports.OperatorRepository) *RegisterRootOperator {
	return &RegisterRootOperator{
		operatorRepository: operatorRepository,
	}
}

func (r *RegisterRootOperator) Execute(
	ctx context.Context,
	cmd *commands.RegisterRootOperator,
) (*aggregates.Operator, error) {
	if existing, _ := r.operatorRepository.FindOneByEmail(ctx, cmd.Email); existing != nil {
		return nil, derrors.ErrOperatorEmailAlreadyExists
	}

	if existing, _ := r.operatorRepository.FindOneByUsername(ctx, cmd.Username); existing != nil {
		return nil, derrors.ErrOperatorUsernameAlreadyExists
	}

	newOperator, errNewOperator := aggregates.NewOperator(
		cmd.Name,
		cmd.Username,
		cmd.Email,
		cmd.PIN,
		valueobjects.RootOperator.ToValue(),
	)

	if errNewOperator != nil {
		return nil, errNewOperator
	}

	if err := r.operatorRepository.Create(ctx, newOperator); err != nil {
		if dbutils.IsUniqueConstraint(err, "operator_email_unique") {
			return nil, derrors.ErrOperatorEmailAlreadyExists
		}
		if dbutils.IsUniqueConstraint(err, "operator_username_unique") {
			return nil, derrors.ErrOperatorUsernameAlreadyExists
		}
		return nil, err
	}

	return newOperator, nil
}
