package usecases

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
)

type SwitchOrganizationUseCase struct {
	organizationRespo ports.OrganizationRepository
	appStateRepo      ports.AppStateRepository
}

func NewSwitchOrganizationUseCase(
	organizationRespo ports.OrganizationRepository,
	appStateRepo ports.AppStateRepository,
) *SwitchOrganizationUseCase {
	return &SwitchOrganizationUseCase{
		organizationRespo: organizationRespo,
		appStateRepo:      appStateRepo,
	}
}

func (uc *SwitchOrganizationUseCase) Execute(
	ctx context.Context,
	cmd *commands.SwitchOrganizationCommand,
) (*aggregates.Organization, error) {
	appState, err := uc.appStateRepo.Get(ctx)

	if err != nil {
		return nil, err
	}

	organization, err := uc.organizationRespo.FindOneByID(ctx, cmd.OrganizationID)

	if err != nil {
		return nil, err
	}

	if organization == nil {
		return nil, derrors.ErrOrganizationNotFound
	}

	appState.SwitchOrganization(organization)

	return organization, nil
}
