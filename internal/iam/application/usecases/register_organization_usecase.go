package usecases

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
	"desktop/internal/iam/domain/valueobjects"
)

type RegisterOrganizationUseCase struct {
	organizationRepo ports.OrganizationRepository
	memberRepo       ports.MemberRepository
}

func NewRegisterOrganizationUseCase(
	organizationRepo ports.OrganizationRepository,
	memberRepo ports.MemberRepository,
) *RegisterOrganizationUseCase {
	return &RegisterOrganizationUseCase{
		organizationRepo: organizationRepo,
		memberRepo:       memberRepo,
	}
}

func (uc *RegisterOrganizationUseCase) Execute(
	ctx context.Context,
	activeOperator aggregates.Operator,
	cmd *commands.RegisterOrganization,
) (*aggregates.Organization, error) {
	if !activeOperator.IsRoot() {
		return nil, derrors.ErrOperatorMustBeRoot
	}

	newOrg, err := aggregates.NewOrganization(
		cmd.Name,
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

	newMember := aggregates.NewMember(newOrg.ID, activeOperator.ID, valueobjects.MemberRoleRoot)
	newOrg.AddMember(newMember)

	err = uc.organizationRepo.Create(ctx, newOrg)

	if err != nil {
		return nil, err
	}

	errMember := uc.memberRepo.Create(ctx, newMember)

	if errMember != nil {
		return nil, errMember
	}

	return newOrg, nil
}
