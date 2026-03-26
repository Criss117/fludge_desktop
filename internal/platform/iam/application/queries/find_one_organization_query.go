package queries

import (
	"context"
	"desktop/internal/platform/iam/application/responses"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/ports"
)

type FindOneOrganization struct {
	organizationRepository ports.OrganizationRepository
}

func NewFindOneOrganization(organizationRepository ports.OrganizationRepository) *FindOneOrganization {
	return &FindOneOrganization{
		organizationRepository: organizationRepository,
	}
}

func (r *FindOneOrganization) Execute(
	ctx context.Context,
	loggedOperatorId string,
	organizationId string,
) (*responses.Organization, error) {
	organization, err := r.organizationRepository.FindOneById(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if organization == nil {
		return nil, derrors.ErrOrganizationNotFound
	}

	if organization.FindMemberByOperatorId(loggedOperatorId) == nil {
		return nil, derrors.ErrMemberNotFound
	}

	orgResponse := responses.OrganizationFromDomain(organization)

	return &orgResponse, nil
}

func (r *FindOneOrganization) ExecuteAggregate(
	ctx context.Context,
	loggedOperatorId string,
	organizationId string,
) (*aggregates.Organization, error) {
	organization, err := r.organizationRepository.FindOneById(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if organization == nil {
		return nil, derrors.ErrOrganizationNotFound
	}

	if organization.FindMemberByOperatorId(loggedOperatorId) == nil {
		return nil, derrors.ErrMemberNotFound
	}

	return organization, nil
}
