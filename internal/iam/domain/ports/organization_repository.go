package ports

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
)

type OrganizationRepository interface {
	// Create(org *aggregates.Organization) error
	// Update(org *aggregates.Organization) error
	// Delete(org *aggregates.Organization) error

	FindOneByID(ctx context.Context, organizationId string) (*aggregates.Organization, error)
	// FindBySlug(slug string) (*aggregates.Organization, error)
	// ExistsBySlug(slug string) (bool, error)
	// ExistsByLegalName(legalName string) (bool, error)
}
