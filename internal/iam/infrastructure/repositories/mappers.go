package repositories

import (
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/iam/infrastructure/dtos"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

func TeamMemberToDomain(dbTeamMember *db.TeamMember) *aggregates.TeamMember {
	return aggregates.ReconstituteTeamMember(
		dbTeamMember.ID,
		dbTeamMember.TeamID,
		dbTeamMember.OperatorID,
		dbTeamMember.OrganizationID,
		platform.FromMillis(dbTeamMember.CreatedAt),
		platform.FromMillis(dbTeamMember.UpdatedAt),
		platform.FromMillisNullable(dbTeamMember.DeletedAt),
	)
}

func TeamToDomain(dbTeam *db.Team, dbTeamMembers []db.TeamMember) (*aggregates.Team, error) {
	domainTeamMembers := make([]*aggregates.TeamMember, len(dbTeamMembers))

	for i, member := range dbTeamMembers {
		domainTeamMembers[i] = TeamMemberToDomain(&member)
	}

	permissions, errPermissions := valueobjects.NewPermissionListFromJSON(dbTeam.Permissions)

	if errPermissions != nil {
		return nil, errPermissions
	}

	return aggregates.ReconstituteTeam(
		dbTeam.ID,
		dbTeam.Name,
		dbTeam.OrganizationID,
		permissions,
		platform.FromStringNullable(dbTeam.Description),
		platform.FromMillis(dbTeam.CreatedAt),
		platform.FromMillis(dbTeam.UpdatedAt),
		platform.FromMillisNullable(dbTeam.DeletedAt),
		domainTeamMembers,
	), nil
}

func MemberToDomain(dbMember *db.Member) *aggregates.Member {
	return aggregates.ReconstituteMember(
		dbMember.ID,
		dbMember.OrganizationID,
		dbMember.OperatorID,
		dbMember.Role,
		platform.FromMillis(dbMember.CreatedAt),
		platform.FromMillis(dbMember.UpdatedAt),
		platform.FromMillisNullable(dbMember.DeletedAt),
	)
}

func OperatorToDomain(dbOperator *db.Operator) *aggregates.Operator {
	return aggregates.ReconstituteOperator(
		dbOperator.ID,
		dbOperator.Name,
		dbOperator.Username,
		dbOperator.Email,
		dbOperator.Pin,
		platform.IntToBool(dbOperator.IsRoot),
		platform.FromMillis(dbOperator.CreatedAt),
		platform.FromMillis(dbOperator.UpdatedAt),
		platform.FromMillisNullable(dbOperator.DeletedAt),
	)
}

func OrganizationToDomain(
	dbActiveOrganization *dtos.DBActiveOrganizationDto,
) *aggregates.Organization {
	domainMembers := make([]*aggregates.Member, len(dbActiveOrganization.Members))

	for i, member := range dbActiveOrganization.Members {
		domainMembers[i] = MemberToDomain(&member)
	}

	return aggregates.ReconstituteOrganization(
		dbActiveOrganization.Organization.ID,
		dbActiveOrganization.Organization.Name,
		dbActiveOrganization.Organization.Slug,
		dbActiveOrganization.Organization.LegalName,
		dbActiveOrganization.Organization.Address,
		platform.FromStringNullable(dbActiveOrganization.Organization.Logo),
		platform.FromStringNullable(dbActiveOrganization.Organization.ContactPhone),
		platform.FromStringNullable(dbActiveOrganization.Organization.ContactEmail),
		platform.FromMillis(dbActiveOrganization.Organization.CreatedAt),
		platform.FromMillis(dbActiveOrganization.Organization.UpdatedAt),
		platform.FromMillisNullable(dbActiveOrganization.Organization.DeletedAt),
		domainMembers,
		nil,
	)
}

func AppStateToDomain(
	dbAppState *db.AppState,
	activeOperator *aggregates.Operator,
) aggregates.AppState {
	var activeOrganizationId *string = nil

	if dbAppState.ActiveOrganizationID.Valid {
		activeOrganizationId = &dbAppState.ActiveOrganizationID.String
	}

	return aggregates.NewAppState(
		activeOrganizationId,
		activeOperator,
		platform.FromMillis(dbAppState.UpdatedAt),
	)
}
