package usecases

import (
	"context"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/platform/iam/domain/valueobjects"
	"desktop/internal/shared/db/dbutils"
)

type RegisterOrganization struct {
	txManager              dbutils.TxManager
	organizationRepository ports.OrganizationRepository
	teamRepository         ports.OrganizationTeamRepository
	memberRepository       ports.OrganizationMemberRepository
}

func NewRegisterOrganization(
	txManager dbutils.TxManager,
	organizationRepository ports.OrganizationRepository,
	teamRepository ports.OrganizationTeamRepository,
	memberRepository ports.OrganizationMemberRepository,
) *RegisterOrganization {
	return &RegisterOrganization{
		txManager:              txManager,
		organizationRepository: organizationRepository,
		teamRepository:         teamRepository,
		memberRepository:       memberRepository,
	}
}

func (r *RegisterOrganization) Execute(
	ctx context.Context,
	loggedOperatorId string,
	cmd *commands.RegisterOrganization,
) (*aggregates.Organization, error) {

	newOrganizations, errNewOrganizations := aggregates.NewOrganization(
		cmd.Name,
		cmd.LegalName,
		cmd.Address,
		cmd.Logo,
		cmd.ContactPhone,
		cmd.ContactEmail,
	)

	if errNewOrganizations != nil {
		return nil, errNewOrganizations
	}

	existsByDetails, errExists := r.organizationRepository.ExistsByDetails(ctx, ports.ExistsByDetails{
		Slug:      newOrganizations.Slug.Value(),
		LegalName: newOrganizations.LegalName,
		Name:      newOrganizations.Name,
	})

	if errExists != nil {
		return nil, errExists
	}

	if existsByDetails > 0 {
		return nil, derrors.ErrOrganizationAlreadyExists
	}

	// errTx := r.txManager.WithTx(ctx, func(q *db.Queries) error {
	// 	if errDb := r.organizationRepository.Create(ctx, newOrganizations); errDb != nil {
	// 		return errDb
	// 	}

	// 	defaultTeam := aggregates.DefaultTeam(newOrganizations.ID)

	// 	if errDb := r.teamRepository.Create(ctx, defaultTeam); errDb != nil {
	// 		return errDb
	// 	}

	// 	return nil
	// })

	// if errTx != nil {
	// 	return nil, errTx
	// }

	if errDb := r.organizationRepository.Create(ctx, newOrganizations); errDb != nil {
		return nil, errDb
	}

	defaultTeam := aggregates.DefaultTeam(newOrganizations.ID)

	if errDb := r.teamRepository.Create(ctx, defaultTeam); errDb != nil {
		return nil, errDb
	}

	rootMember := aggregates.NewMember(
		newOrganizations.ID,
		loggedOperatorId,
		valueobjects.MemberRoleRoot,
	)

	if errDb := r.memberRepository.Create(ctx, rootMember); errDb != nil {
		return nil, errDb
	}

	return newOrganizations, nil
}
