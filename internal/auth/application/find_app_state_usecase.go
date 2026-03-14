package application

import (
	"desktop/internal/auth/application/dtos"
	"desktop/internal/auth/domain"
)

type FindAppStateUseCase struct {
	appStateRepository domain.AppStateRepository
}

func NewFindAppStateUseCase(appStateRepository domain.AppStateRepository) *FindAppStateUseCase {
	return &FindAppStateUseCase{
		appStateRepository: appStateRepository,
	}
}

func (uc *FindAppStateUseCase) Execute() (*dtos.AppStateDTO, error) {
	appState, err := uc.appStateRepository.FindAppState()

	if err != nil {
		return nil, err
	}

	return dtos.DomainToAppStateDTO(appState), nil
}

func (uc *FindAppStateUseCase) ExecuteWithoutParsing() (*domain.AppState, error) {
	appState, err := uc.appStateRepository.FindAppState()

	if err != nil {
		return nil, err
	}

	return appState, nil
}
