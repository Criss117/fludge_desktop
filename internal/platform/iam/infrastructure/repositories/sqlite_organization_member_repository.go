package repositories

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
	"desktop/internal/shared/derrors"
)

type SqliteOrganizationMemberRepository struct {
	queries *db.Queries
}

func NewSqliteOrganizationMemberRepository(queries *db.Queries) *SqliteOrganizationMemberRepository {
	return &SqliteOrganizationMemberRepository{
		queries: queries,
	}
}

func (r *SqliteOrganizationMemberRepository) Create(ctx context.Context, member *aggregates.Member) error {
	if errDb := r.queries.CreateMember(ctx, db.CreateMemberParams{
		ID:             member.ID,
		OrganizationID: member.OrganizationID,
		OperatorID:     member.OperatorID,
		Role:           member.Role.Value(),
		CreatedAt:      dbutils.TimeToInt64(member.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(member.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOrganizationMemberRepository) Delete(ctx context.Context, member *aggregates.Member) error {
	if member.DeletedAt != nil {
		return derrors.ErrInvalidDeleteState
	}

	if errDb := r.queries.DeleteMember(ctx, db.DeleteMemberParams{
		ID:        member.ID,
		DeletedAt: dbutils.TimeToSQLNullable(member.DeletedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOrganizationMemberRepository) FindAllByOrganization(ctx context.Context, organizationId string) ([]*aggregates.Member, error) {
	dbMembers, err := r.queries.FindAllMembers(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbMembers) == 0 {
		return nil, nil
	}

	members := make([]*aggregates.Member, len(dbMembers))

	for i, dbMember := range dbMembers {
		members[i] = mappers.MapMemberToDomain(dbMember)
	}

	return members, nil
}
