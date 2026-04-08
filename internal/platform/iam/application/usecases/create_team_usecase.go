package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/ports"
)

type CreateTeam struct {
	organizaionTeamRepository ports.OrganizationTeamRepository
}

func NewCreateTeam(organizaionTeamRepository ports.OrganizationTeamRepository) *CreateTeam {
	return &CreateTeam{
		organizaionTeamRepository: organizaionTeamRepository,
	}
}

func (uc *CreateTeam) Execute(
	ctx context.Context,
	currentOrganization *aggregates.Organization,
	cmd *commands.CreateTeam,
) (*aggregates.Team, error) {
	newTeam, err := currentOrganization.NewTeam(
		cmd.Name,
		cmd.Description,
		cmd.Permissions,
	)

	if err != nil {
		return nil, err
	}

	if err := uc.organizaionTeamRepository.Create(ctx, newTeam); err != nil {
		return nil, err
	}

	return newTeam, nil
}
