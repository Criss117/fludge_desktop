package ports

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
)

type FindManyOrganizationsBy struct {
	Slug      string
	LegalName string
	Name      string
}

type ExistsByDetails struct {
	Slug      string
	LegalName string
	Name      string
}

type OrganizationRepository interface {
	Create(ctx context.Context, organization *aggregates.Organization) error
	Update(ctx context.Context, organization *aggregates.Organization) error

	ExistsByDetails(ctx context.Context, details ExistsByDetails) (int64, error)

	FindOneById(ctx context.Context, organizationId string) (*aggregates.Organization, error)
	FindManyByRootOperator(ctx context.Context, operatorId string) ([]*aggregates.Organization, error)
}

type OrganizationMemberRepository interface {
	Create(ctx context.Context, member *aggregates.Member) error
	Delete(ctx context.Context, member *aggregates.Member) error

	FindAllByOrganization(ctx context.Context, organizationId string) ([]*aggregates.Member, error)
}

type OrganizationTeamRepository interface {
	Create(ctx context.Context, team *aggregates.Team) error
	Delete(ctx context.Context, team *aggregates.Team) error

	FindAllByOrganization(ctx context.Context, organizationId string) ([]*aggregates.Team, error)
}
