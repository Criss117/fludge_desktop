package mappers

import (
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/valueobjects"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type DBTeam struct {
	db.Team
	Members []db.TeamMember
}

func MapMemberFromDomain(member *aggregates.Member) db.Member {
	if member == nil {
		return db.Member{}
	}

	return db.Member{
		ID:             member.ID,
		OrganizationID: member.OrganizationID,
		OperatorID:     member.OperatorID,
		Role:           member.Role.Value(),
		CreatedAt:      dbutils.TimeToInt64(member.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(member.UpdatedAt),
		DeletedAt:      dbutils.TimeToSQLNullable(member.DeletedAt),
	}
}

func MapTeamFromDomain(team *aggregates.Team) (db.Team, error) {
	if team == nil {
		return db.Team{}, nil
	}

	permissions, err := team.Permissions.ToJSON()

	if err != nil {
		return db.Team{}, err
	}

	return db.Team{
		ID:             team.ID,
		Name:           team.Name,
		OrganizationID: team.OrganizationID,
		Permissions:    permissions,
		Description:    dbutils.StringToSQLNullable(team.Description),
		CreatedAt:      dbutils.TimeToInt64(team.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(team.UpdatedAt),
		DeletedAt:      dbutils.TimeToSQLNullable(team.DeletedAt),
	}, nil
}

func MapTeamMemberFromDomain(member *aggregates.TeamMember) db.TeamMember {
	if member == nil {
		return db.TeamMember{}
	}

	return db.TeamMember{
		ID:             member.ID,
		TeamID:         member.TeamID,
		OperatorID:     member.OperatorID,
		OrganizationID: member.OrganizationID,
		CreatedAt:      dbutils.TimeToInt64(member.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(member.UpdatedAt),
		DeletedAt:      dbutils.TimeToSQLNullable(member.DeletedAt),
	}
}

func MapOrganizationFromDomain(organization *aggregates.Organization) db.Organization {
	if organization == nil {
		return db.Organization{}
	}

	var cemail *string = nil

	if organization.ContactEmail != nil {
		email := organization.ContactEmail.Value()
		cemail = &email
	}

	return db.Organization{
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
		DeletedAt:    dbutils.TimeToSQLNullable(organization.DeletedAt),
	}
}

func MapMemberToDomain(member db.Member) *aggregates.Member {
	return aggregates.ReconstituteMember(
		member.ID,
		member.OrganizationID,
		member.OperatorID,
		member.Role,
		dbutils.TimeFromInt64(member.CreatedAt),
		dbutils.TimeFromInt64(member.UpdatedAt),
		dbutils.TimeFromSQLNullable(member.DeletedAt),
	)
}

func MapTeamMemberToDomain(member db.TeamMember) *aggregates.TeamMember {
	return aggregates.ReconstituteTeamMember(
		member.ID,
		member.TeamID,
		member.OperatorID,
		member.OrganizationID,
		dbutils.TimeFromInt64(member.CreatedAt),
		dbutils.TimeFromInt64(member.UpdatedAt),
		dbutils.TimeFromSQLNullable(member.DeletedAt),
	)
}

func MapTeamToDomain(team DBTeam) *aggregates.Team {
	permissions, err := valueobjects.NewPermissionListFromJSON(team.Permissions)

	if err != nil {
		return nil
	}

	dMembers := make([]*aggregates.TeamMember, len(team.Members))

	for i, member := range team.Members {
		dMembers[i] = MapTeamMemberToDomain(member)
	}

	return aggregates.ReconstituteTeam(
		team.ID,
		team.Name,
		team.OrganizationID,
		permissions,
		dbutils.StringFromSQLNullable(team.Description),
		dbutils.TimeFromInt64(team.CreatedAt),
		dbutils.TimeFromInt64(team.UpdatedAt),
		dbutils.TimeFromSQLNullable(team.DeletedAt),
		dMembers,
	)
}

func MapOrganizationToDomain(organization db.Organization, members []*aggregates.Member, teams []*aggregates.Team) *aggregates.Organization {
	org := aggregates.ReconstituteOrganization(
		organization.ID,
		organization.Name,
		organization.Slug,
		organization.LegalName,
		organization.Address,
		dbutils.StringFromSQLNullable(organization.Logo),
		dbutils.StringFromSQLNullable(organization.ContactPhone),
		dbutils.StringFromSQLNullable(organization.ContactEmail),
		dbutils.TimeFromInt64(organization.CreatedAt),
		dbutils.TimeFromInt64(organization.UpdatedAt),
		dbutils.TimeFromSQLNullable(organization.DeletedAt),
		members,
		teams,
	)

	return org
}
