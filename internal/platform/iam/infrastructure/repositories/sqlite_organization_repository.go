package repositories

import (
	"context"
	"database/sql"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/platform/iam/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
	"errors"
)

type SqliteOrganizationRepository struct {
	queries          *db.Queries
	memberRepository ports.OrganizationMemberRepository
	teamRepository   ports.OrganizationTeamRepository
}

func NewSqliteOrganizationRepository(queries *db.Queries) *SqliteOrganizationRepository {
	return &SqliteOrganizationRepository{
		queries: queries,
	}
}

func (r *SqliteOrganizationRepository) Create(ctx context.Context, organization *aggregates.Organization) error {
	exists, errExists := r.queries.ExistsOrganization(ctx, db.ExistsOrganizationParams{
		Name:      organization.Name,
		LegalName: organization.LegalName,
		Slug:      organization.Slug.Value(),
	})

	if errExists != nil {
		return errExists
	}

	if exists > 0 {
		return derrors.ErrOrganizationAlreadyExists
	}

	var cemail *string = nil

	if organization.ContactEmail != nil {
		email := organization.ContactEmail.Value()
		cemail = &email
	}

	if errDb := r.queries.CreateOrganization(ctx, db.CreateOrganizationParams{
		ID:           organization.ID,
		Name:         organization.Name,
		Slug:         organization.Slug.Value(),
		LegalName:    organization.LegalName,
		Address:      organization.Address,
		Logo:         dbutils.StringToSQLNullable(organization.Logo),
		ContactPhone: dbutils.StringToSQLNullable(organization.ContactPhone),
		ContactEmail: dbutils.StringToSQLNullable(cemail),
		CreatedAt:    dbutils.TimeToInt64(organization.CreatedAt),
		UpdatedAt:    dbutils.TimeToInt64(organization.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOrganizationRepository) Update(ctx context.Context, organization *aggregates.Organization) error {
	var cemail *string = nil

	if organization.ContactEmail != nil {
		email := organization.ContactEmail.Value()
		cemail = &email
	}

	if err := r.queries.UpdateOrganization(ctx, db.UpdateOrganizationParams{
		ID:           organization.ID,
		Name:         organization.Name,
		Slug:         organization.Slug.Value(),
		LegalName:    organization.LegalName,
		Address:      organization.Address,
		Logo:         dbutils.StringToSQLNullable(organization.Logo),
		ContactPhone: dbutils.StringToSQLNullable(organization.ContactPhone),
		ContactEmail: dbutils.StringToSQLNullable(cemail),
		UpdatedAt:    dbutils.TimeToInt64(organization.UpdatedAt),
	}); err != nil {
		return err
	}

	return nil
}

func (r *SqliteOrganizationRepository) FindOneById(ctx context.Context, organizationId string) (*aggregates.Organization, error) {
	organization, err := r.queries.FindOneOrganization(ctx, organizationId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	members, errMembers := r.memberRepository.FindAllByOrganization(ctx, organizationId)

	if errMembers != nil {
		return nil, errMembers
	}

	teams, errTeams := r.teamRepository.FindAllByOrganization(ctx, organizationId)

	if errTeams != nil {
		return nil, errTeams
	}

	return mappers.MapOrganizationToDomain(organization, members, teams), nil
}

func (r *SqliteOrganizationRepository) FindManyByRootOperator(ctx context.Context, operatorId string) ([]*aggregates.Organization, error) {
	dbOrganizations, err := r.queries.FindManyOrganizationsByRootOperator(ctx, operatorId)

	if err != nil {
		return nil, err
	}

	if len(dbOrganizations) == 0 {
		return nil, nil
	}

	organizations := make([]*aggregates.Organization, len(dbOrganizations))

	for i, dbOrganization := range dbOrganizations {
		members, errMembers := r.memberRepository.FindAllByOrganization(ctx, dbOrganization.ID)

		if errMembers != nil {
			return nil, errMembers
		}

		teams, errTeams := r.teamRepository.FindAllByOrganization(ctx, dbOrganization.ID)

		if errTeams != nil {
			return nil, errTeams
		}

		organizations[i] = mappers.MapOrganizationToDomain(dbOrganization, members, teams)
	}

	return organizations, nil
}

func (r *SqliteOrganizationRepository) ExistsByDetails(ctx context.Context, details ports.ExistsByDetails) (int64, error) {
	exists, err := r.queries.ExistsOrganization(ctx, db.ExistsOrganizationParams{
		Name:      details.Name,
		LegalName: details.LegalName,
		Slug:      details.Slug,
	})

	if err != nil {
		return 0, err
	}

	return exists, nil
}
