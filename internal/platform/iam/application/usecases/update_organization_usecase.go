package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/ports"
)

type UpdateOrganization struct {
	OrganizationRepository ports.OrganizationRepository
}

func NewUpdateOrganization(organizationRepository ports.OrganizationRepository) *UpdateOrganization {
	return &UpdateOrganization{
		OrganizationRepository: organizationRepository,
	}
}

func (u *UpdateOrganization) Execute(
	ctx context.Context,
	loggedOperatorId string,
	cmd *commands.UpdateOrganization,
) (*aggregates.Organization, error) {
	organization, err := u.OrganizationRepository.FindOneById(ctx, cmd.ID)

	if err != nil {
		return nil, err
	}

	if member := organization.FindMemberByOperatorId(loggedOperatorId); member == nil {
		return nil, derrors.ErrOperatorIsNotMemberOfOrg
	}

	organization.UpdateDetails(
		cmd.Name,
		cmd.LegalName,
		cmd.Address,
		cmd.Logo,
		cmd.ContactPhone,
		cmd.ContactEmail,
	)

	if err := u.OrganizationRepository.Update(ctx, organization); err != nil {
		return nil, err
	}

	return organization, nil
}
