package usecases

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
)

type RegisterOrganizationUseCase struct {
	organizationRepo ports.OrganizationRepository
}

func NewRegisterOrganizationUseCase(organizationRepo ports.OrganizationRepository) *RegisterOrganizationUseCase {
	return &RegisterOrganizationUseCase{
		organizationRepo: organizationRepo,
	}
}

func (uc *RegisterOrganizationUseCase) Execute(
	ctx context.Context,
	cmd *commands.RegisterOrganizationCommand,
) (*aggregates.Organization, error) {
	newOrg, err := aggregates.NewOrganization(
		cmd.Name,
		cmd.Slug,
		cmd.LegalName,
		cmd.Address,
		cmd.Logo,
		cmd.ContactPhone,
		cmd.ContactEmail,
	)

	if err != nil {
		return nil, err
	}

	exisitingOrganization, err := uc.organizationRepo.FindManyOrganizationsBy(ctx,
		ports.FindManyOrganizationsBy{
			Slug:      newOrg.Slug.Value(),
			LegalName: newOrg.LegalName,
			Name:      newOrg.Name,
		})

	if err != nil {
		return nil, err
	}

	if len(exisitingOrganization) > 0 {
		return nil, derrors.ErrOrganizationAlreadyExists
	}

	err = uc.organizationRepo.Create(ctx, newOrg)

	if err != nil {
		return nil, err
	}

	return newOrg, nil
}
