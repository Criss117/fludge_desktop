package usecases

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
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
) ([]*aggregates.Organization, error) {
	return u.organizationRepository.FindManyByRootOperator(ctx, loggedOperatorId)
}
