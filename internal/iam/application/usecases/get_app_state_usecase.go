package usecases

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/ports"
)

type GetAppStateUseCase struct {
	appStateRepo     ports.AppStateRepository
	organizationRepo ports.OrganizationRepository
}

type GetAppStateResponse struct {
	ActiveOpertaror    *aggregates.Operator
	ActiveOrganization *aggregates.Organization
}

func NewGetAppStateUseCase(
	appStateRepo ports.AppStateRepository,
	organizationRepo ports.OrganizationRepository,
) *GetAppStateUseCase {
	return &GetAppStateUseCase{
		appStateRepo:     appStateRepo,
		organizationRepo: organizationRepo,
	}
}

func (uc *GetAppStateUseCase) Execute(ctx context.Context) (*aggregates.AppState, error) {
	appState, err := uc.appStateRepo.Get(ctx)

	if err != nil {
		return nil, err
	}

	return appState, nil
}
