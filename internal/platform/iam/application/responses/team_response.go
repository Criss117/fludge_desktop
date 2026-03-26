package responses

import (
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/shared/db/dbutils"
)

type TeamMember struct {
	ID             string `json:"id"`
	TeamID         string `json:"teamId"`
	OperatorID     string `json:"operatorId"`
	OrganizationID string `json:"organizationId"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	DeletedAt      *int64 `json:"deletedAt"`
}

type Team struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	OrganizationId string   `json:"organizationId"`
	Permissions    []string `json:"permissions"`
	Description    *string  `json:"description"`
	CreatedAt      int64    `json:"createdAt"`
	UpdatedAt      int64    `json:"updatedAt"`
	DeletedAt      *int64   `json:"deletedAt"`
	Members        []TeamMember
}

func TeamMemberFromDomain(teamMember *aggregates.TeamMember) *TeamMember {
	return &TeamMember{
		ID:             teamMember.ID,
		TeamID:         teamMember.TeamID,
		OperatorID:     teamMember.OperatorID,
		OrganizationID: teamMember.OrganizationID,
		CreatedAt:      dbutils.TimeToInt64(teamMember.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(teamMember.UpdatedAt),
		DeletedAt:      dbutils.TimeToInt64Nullable(teamMember.DeletedAt),
	}
}

func TeamFromDomain(team *aggregates.Team) *Team {

	teamsMembers := make([]TeamMember, len(team.Members))

	for i, member := range team.Members {
		tm := TeamMemberFromDomain(member)

		teamsMembers[i] = *tm
	}

	permissions := make([]string, len(team.Permissions))

	for i, p := range team.Permissions {
		permissions[i] = p.Value()
	}

	return &Team{
		ID:             team.ID,
		Name:           team.Name,
		OrganizationId: team.OrganizationID,
		Permissions:    permissions,
		Description:    team.Description,
		CreatedAt:      dbutils.TimeToInt64(team.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(team.UpdatedAt),
		DeletedAt:      dbutils.TimeToInt64Nullable(team.DeletedAt),
		Members:        teamsMembers,
	}

}
