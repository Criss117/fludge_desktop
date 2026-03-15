package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/infrastructure/dtos"
	"desktop/internal/shared/db"
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
	dbActiveOrganizations, err := sql.queries.FindOneOrganizationByID(ctx, organizationId)

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
