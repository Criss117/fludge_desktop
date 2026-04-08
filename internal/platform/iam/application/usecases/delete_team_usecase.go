package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/ports"
)

type DeleteTeam struct {
	organizationTeamRepository ports.OrganizationTeamRepository
}

func NewDeleteTeam(
	organizationTeamRepository ports.OrganizationTeamRepository,
) *DeleteTeam {
	return &DeleteTeam{
		organizationTeamRepository: organizationTeamRepository,
	}
}

func (uc *DeleteTeam) Execute(
	ctx context.Context,
	currentOrganization *aggregates.Organization,
	cmd commands.DeleteTeam,
) error {
	if err := currentOrganization.RemoveTeam(cmd.ID); err != nil {
		return err
	}

	if err := uc.organizationTeamRepository.Delete(ctx, currentOrganization.ID, cmd.ID); err != nil {
		return err
	}

	return nil
}
