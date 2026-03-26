package queries

import (
	"context"
	"desktop/internal/platform/iam/application/responses"
	"desktop/internal/platform/iam/domain/ports"
)

type FindManyOrganizationsByRootOperator struct {
	organizationRepository ports.OrganizationRepository
}

func NewFindManyOrganizationsByRootOperator(
	organizationRepository ports.OrganizationRepository,
) *FindManyOrganizationsByRootOperator {
	return &FindManyOrganizationsByRootOperator{
		organizationRepository: organizationRepository,
	}
}

func (u *FindManyOrganizationsByRootOperator) Execute(
	ctx context.Context,
	loggedOperatorId string,
) ([]responses.Organization, error) {
	organizations, err := u.organizationRepository.FindManyByRootOperator(ctx, loggedOperatorId)

	if err != nil {
		return nil, err
	}

	if len(organizations) == 0 {
		return make([]responses.Organization, 0), nil
	}

	orgs := make([]responses.Organization, len(organizations))

	for i, org := range organizations {
		o := responses.OrganizationFromDomain(org)

		orgs[i] = o
	}

	return orgs, nil
}
