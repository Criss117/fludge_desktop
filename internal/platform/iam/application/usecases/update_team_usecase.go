package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/ports"
)

type UpdateTeam struct {
	organizationTeamRepository ports.OrganizationTeamRepository
}

func NewUpdateTeam(
	organizationTeamRepository ports.OrganizationTeamRepository,
) *UpdateTeam {
	return &UpdateTeam{
		organizationTeamRepository: organizationTeamRepository,
	}
}

func (uc *UpdateTeam) Execute(
	ctx context.Context,
	currentOrganization *aggregates.Organization,
	cmd commands.UpdateTeam,
) error {
	teamToUpdate, err := currentOrganization.UpdateTeam(
		cmd.ID,
		cmd.Name,
		cmd.Description,
		cmd.Permissions,
	)

	if err != nil {
		return err
	}

	if err := uc.organizationTeamRepository.Update(ctx, teamToUpdate); err != nil {
		return err
	}

	return nil
}
