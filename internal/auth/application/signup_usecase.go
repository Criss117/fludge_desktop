package application

import (
	"desktop/internal/auth/application/dtos"
	"desktop/internal/auth/domain"
	"desktop/internal/shared/dberrors"
)

type SignUpUseCase struct {
	operatorRepository domain.OperatorRepository
	appStateRepository domain.AppStateRepository
}

func NewSignUpUseCase(operatorRepository domain.OperatorRepository, appStateRepository domain.AppStateRepository) *SignUpUseCase {
	return &SignUpUseCase{
		operatorRepository: operatorRepository,
		appStateRepository: appStateRepository,
	}
}

func (uc *SignUpUseCase) Execute(operatorToCreate dtos.SignupDTO) (*domain.AppState, error) {
	existingOperators, err := uc.operatorRepository.FindManyByUsernameOrEmail(
		operatorToCreate.Username,
		operatorToCreate.Email,
	)

	if err != nil {
		return nil, dberrors.ErrDatabaseQuery()
	}

	for _, op := range existingOperators {
		if op.Username == operatorToCreate.Username {
			return nil, domain.ErrUsernameAlreadyExists()
		}
		if op.Email == operatorToCreate.Email {
			return nil, domain.ErrEmailAlreadyExists()
		}
	}

	newOperator := domain.NewOperator(
		operatorToCreate.Name,
		operatorToCreate.Username,
		operatorToCreate.Email,
		operatorToCreate.Pin,
		true,
	)

	errCreating := uc.operatorRepository.Create(newOperator)

	if errCreating != nil {
		return nil, domain.ErrOperatorNotCreated()
	}

	currentAppState, errAppState := uc.appStateRepository.FindAppState()

	if errAppState != nil {
		return nil, errAppState
	}

	currentAppState.SetOperator(newOperator)

	errUpdating := uc.appStateRepository.Update(*currentAppState)

	if errUpdating != nil {
		return nil, errUpdating
	}

	return currentAppState, nil
}
