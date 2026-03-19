package ports

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
)

type FindManyOrganizationsBy struct {
	Slug      string
	LegalName string
	Name      string
}

type OrganizationRepository interface {
	Create(ctx context.Context, organization *aggregates.Organization) error
	// Update(org *aggregates.Organization) error
	// Delete(org *aggregates.Organization) error

	FindOneByID(ctx context.Context, organizationId string) (*aggregates.Organization, error)
	FindByOperator(ctx context.Context, operatorId string) ([]*aggregates.Organization, error)
	FindManyOrganizationsBy(ctx context.Context, values FindManyOrganizationsBy) ([]*aggregates.Organization, error)
}
