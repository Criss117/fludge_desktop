package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db"
)

type SqliteTeamRepository struct {
	queries *db.Queries
}

func NewSqliteTeamRepository(queries *db.Queries) *SqliteTeamRepository {
	return &SqliteTeamRepository{
		queries: queries,
	}
}

func (sql *SqliteTeamRepository) FindAllByOperatorId(ctx context.Context, operatorId string) ([]*aggregates.Team, error) {
	dbTeams, err := sql.queries.FindAllTeamsByOperatorId(ctx, operatorId)

	if err != nil {
		return nil, err
	}

	if len(dbTeams) == 0 {
		return nil, nil
	}

	teams := make([]*aggregates.Team, len(dbTeams))

	for i, dbTeam := range dbTeams {
		dbTeamMembers, err := sql.queries.FindAllTeamsMembersByTeamId(ctx, dbTeam.ID)

		if err != nil {
			return nil, err
		}

		team, errTeam := TeamToDomain(&dbTeam, dbTeamMembers)

		if errTeam != nil {
			return nil, errTeam
		}

		teams[i] = team
	}

	return teams, nil
}
