package usecases

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/ports"
)

type FindOpertaroOrganizations struct {
	organizationRepo ports.OrganizationRepository
}

func NewFindOpertaroOrganizations(organizationRepo ports.OrganizationRepository) *FindOpertaroOrganizations {
	return &FindOpertaroOrganizations{
		organizationRepo: organizationRepo,
	}
}

func (uc *FindOpertaroOrganizations) Execute(ctx context.Context, operatorId string) ([]*aggregates.Organization, error) {
	return uc.organizationRepo.FindByOperator(ctx, operatorId)
}
