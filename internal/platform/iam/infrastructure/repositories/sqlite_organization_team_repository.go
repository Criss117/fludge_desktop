package repositories

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type SqliteOrganizationTeamRepository struct {
	queries *db.Queries
}

func NewSqliteOrganizationTeamRepository(queries *db.Queries) *SqliteOrganizationTeamRepository {
	return &SqliteOrganizationTeamRepository{
		queries: queries,
	}
}

func (r *SqliteOrganizationTeamRepository) Create(ctx context.Context, team *aggregates.Team) error {
	exists, errExists := r.queries.ExistsTeam(ctx, db.ExistsTeamParams{
		Name:           team.Name,
		OrganizationID: team.OrganizationID,
	})

	if errExists != nil {
		return errExists
	}

	if exists > 0 {
		return derrors.ErrTeamAlreadyExists
	}

	permissions, err := team.Permissions.ToJSON()

	if err != nil {
		return err
	}

	if errDb := r.queries.CreateTeam(ctx, db.CreateTeamParams{
		ID:             team.ID,
		Name:           team.Name,
		OrganizationID: team.OrganizationID,
		Permissions:    permissions,
		Description:    dbutils.StringToSQLNullable(team.Description),
		CreatedAt:      dbutils.TimeToInt64(team.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(team.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOrganizationTeamRepository) Delete(ctx context.Context, team *aggregates.Team) error {
	if err := r.queries.DeleteTeam(ctx, team.ID); err != nil {
		return err
	}

	return nil
}

func (r *SqliteOrganizationTeamRepository) FindAllByOrganization(ctx context.Context, organizationId string) ([]*aggregates.Team, error) {
	dbTeams, err := r.queries.FindAllTeams(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbTeams) == 0 {
		return nil, nil
	}

	teamsWithMembers := make([]mappers.DBTeam, len(dbTeams))

	for i, team := range dbTeams {
		members, errMembers := r.queries.FindAllTeamsMembersByTeam(ctx, team.ID)

		if errMembers != nil {
			return nil, errMembers
		}

		teamsWithMembers[i] = mappers.DBTeam{
			Team:    team,
			Members: members,
		}
	}

	teams := make([]*aggregates.Team, len(teamsWithMembers))

	for i, team := range teamsWithMembers {
		teams[i] = mappers.MapTeamToDomain(team)
	}

	return teams, nil
}
