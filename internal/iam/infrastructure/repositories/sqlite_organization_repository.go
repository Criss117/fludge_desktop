package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/ports"
	"desktop/internal/iam/infrastructure/dtos"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SqliteOrganizationRepository struct {
	queries *db.Queries
}

func NewSqliteOrganizationRepository(queries *db.Queries) *SqliteOrganizationRepository {
	return &SqliteOrganizationRepository{
		queries: queries,
	}
}

func (sql *SqliteOrganizationRepository) FindOneByID(ctx context.Context, organizationId string) (*aggregates.Organization, error) {
	dbActiveOrganizations, err := sql.queries.FindOneOrganizationById(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbActiveOrganizations) == 0 {
		return nil, derrors.ErrOrganizationNotFound
	}

	currentOrganization := dbActiveOrganizations[0]

	dbOrganizationMembers, err := sql.queries.FindAllMembersByOrganizationId(ctx, currentOrganization.ID)

	if err != nil {
		return nil, err
	}

	dbOrganizationTeams, err := sql.queries.FindAllTeamsByOrganizationId(ctx, currentOrganization.ID)

	if err != nil {
		return nil, err
	}

	dbOrganizationTeamsMembers := make([]dtos.DBActiveTeamDto, len(dbOrganizationTeams))

	for i, team := range dbOrganizationTeams {
		dbTeamMembers, err := sql.queries.FindAllTeamsMembersByTeamId(ctx, team.ID)

		if err != nil {
			return nil, err
		}

		dbOrganizationTeamsMembers[i] = dtos.DBActiveTeamDto{
			Members: dbTeamMembers,
			Team:    team,
		}
	}

	dbActiveOrganization := &dtos.DBActiveOrganizationDto{
		Organization: &currentOrganization,
		Members:      dbOrganizationMembers,
		Teams:        dbOrganizationTeamsMembers,
	}

	return OrganizationToDomain(dbActiveOrganization), nil
}

func (sql *SqliteOrganizationRepository) FindByOperator(ctx context.Context, operatorId string) ([]*aggregates.Organization, error) {
	dbOrganizations, err := sql.queries.FindManyOrganizationsByOperatorId(ctx, operatorId)

	if err != nil {
		return nil, err
	}

	organizations := make([]*aggregates.Organization, len(dbOrganizations))

	for index, dbOrganization := range dbOrganizations {
		dbActiveOrganization := &dtos.DBActiveOrganizationDto{
			Organization: &dbOrganization,
		}

		organizations[index] = OrganizationToDomain(dbActiveOrganization)
	}

	return organizations, nil
}

func (sql *SqliteOrganizationRepository) FindManyOrganizationsBy(
	ctx context.Context,
	values ports.FindManyOrganizationsBy,
) ([]*aggregates.Organization, error) {
	dbOrganizations, err := sql.queries.FindManyOrganizationsBy(ctx, db.FindManyOrganizationsByParams{
		Slug:      values.Slug,
		LegalName: values.LegalName,
		Name:      values.Name,
	})

	if err != nil {
		return nil, err
	}

	organizations := make([]*aggregates.Organization, len(dbOrganizations))

	for index, dbOrganization := range dbOrganizations {
		dbActiveOrganization := &dtos.DBActiveOrganizationDto{
			Organization: &dbOrganization,
		}

		organizations[index] = OrganizationToDomain(dbActiveOrganization)
	}

	return organizations, nil
}

func (sql *SqliteOrganizationRepository) Create(ctx context.Context, organization *aggregates.Organization) error {

	var contactEmail *string = nil

	if organization.ContactEmail != nil {
		ce := organization.ContactEmail.Value()
		contactEmail = &ce
	}

	errCreate := sql.queries.CreateOrganization(ctx, db.CreateOrganizationParams{
		ID:           organization.ID,
		Name:         organization.Name,
		Slug:         organization.Slug.Value(),
		LegalName:    organization.LegalName,
		Address:      organization.Address,
		Logo:         platform.ToStringNullable(organization.Logo),
		ContactPhone: platform.ToStringNullable(organization.ContactPhone),
		ContactEmail: platform.ToStringNullable(contactEmail),
		CreatedAt:    platform.ToMillis(organization.CreatedAt),
		UpdatedAt:    platform.ToMillis(organization.UpdatedAt),
	})

	if errCreate != nil {
		return errCreate
	}

	return nil
}
