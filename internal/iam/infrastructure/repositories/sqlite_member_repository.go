package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SqliteMemberRepository struct {
	queries *db.Queries
}

func NewSqliteMemberRepository(queries *db.Queries) *SqliteMemberRepository {
	return &SqliteMemberRepository{
		queries: queries,
	}
}

func (sql *SqliteMemberRepository) Create(ctx context.Context, member *aggregates.Member) error {
	err := sql.queries.CreateMember(ctx, db.CreateMemberParams{
		ID:             member.ID,
		OrganizationID: member.OrganizationID,
		OperatorID:     member.OperatorID,
		Role:           member.Role.Value(),
		CreatedAt:      platform.ToMillis(member.CreatedAt),
		UpdatedAt:      platform.ToMillis(member.UpdatedAt),
	})

	if err != nil {
		return err
	}

	return nil
}
