package application

import (
	"desktop/internal/auth/domain"
)

type SignInUseCase struct {
	operatorRepository domain.OperatorRepository
	appStateRepository domain.AppStateRepository
}

func NewSignInUseCase(operatorRepository domain.OperatorRepository, appStateRepository domain.AppStateRepository) *SignInUseCase {
	return &SignInUseCase{
		operatorRepository: operatorRepository,
		appStateRepository: appStateRepository,
	}
}

func (uc *SignInUseCase) Execute(username string, pin string) (*domain.AppState, error) {
	existingOperator, errExisting := uc.operatorRepository.FindOneByUsername(username)

	if errExisting != nil {
		return nil, domain.ErrOperatorNotFoundByUsername()
	}

	if existingOperator == nil {
		return nil, domain.ErrOperatorNotFoundByUsername()
	}

	if existingOperator.Pin != pin {
		return nil, domain.ErrInvalidCredentials()
	}

	currentAppState, errAppState := uc.appStateRepository.FindAppState()

	if errAppState != nil {
		return nil, errAppState
	}

	currentAppState.SetOperator(*existingOperator)

	errUpdating := uc.appStateRepository.Update(*currentAppState)

	if errUpdating != nil {
		return nil, errUpdating
	}

	return currentAppState, nil
}
