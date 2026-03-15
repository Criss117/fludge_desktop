package responses

import (
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type TeamMemberResponse struct {
	ID             string `json:"id"`
	TeamID         string `json:"teamId"`
	OperatorID     string `json:"operatorId"`
	OrganizationID string `json:"organizationId"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	DeletedAt      *int64 `json:"deletedAt"`
}

type TeamResponse struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	OrganizationId string   `json:"organizationId"`
	Permissions    []string `json:"permissions"`
	Description    *string  `json:"description"`
	CreatedAt      int64    `json:"createdAt"`
	UpdatedAt      int64    `json:"updatedAt"`
	DeletedAt      *int64   `json:"deletedAt"`
	Members        []TeamMemberResponse
}

func TeamMemberResponseFromDomain(member *aggregates.TeamMember) *TeamMemberResponse {
	primitiveTeamMember := member.ToValues()

	return &TeamMemberResponse{
		ID:             primitiveTeamMember.ID,
		TeamID:         primitiveTeamMember.TeamID,
		OperatorID:     primitiveTeamMember.OperatorID,
		OrganizationID: primitiveTeamMember.OrganizationID,
		CreatedAt:      platform.TimeToInt64(primitiveTeamMember.CreatedAt),
		UpdatedAt:      platform.TimeToInt64(primitiveTeamMember.UpdatedAt),
		DeletedAt:      platform.TimeToInt64Nullable(primitiveTeamMember.DeletedAt),
	}
}

func TeamResponseFromDomain(team *aggregates.Team) *TeamResponse {
	primitiveTeam := team.ToValues()

	teamsMembers := make([]TeamMemberResponse, len(primitiveTeam.Members))

	for i, member := range team.Members() {
		tm := TeamMemberResponseFromDomain(member)

		teamsMembers[i] = *tm
	}

	return &TeamResponse{
		ID:             primitiveTeam.ID,
		Name:           primitiveTeam.Name,
		OrganizationId: primitiveTeam.OrganizationID,
		Permissions:    primitiveTeam.Permissions,
		Description:    primitiveTeam.Description,
		CreatedAt:      platform.TimeToInt64(primitiveTeam.CreatedAt),
		UpdatedAt:      platform.TimeToInt64(primitiveTeam.UpdatedAt),
		DeletedAt:      platform.TimeToInt64Nullable(primitiveTeam.DeletedAt),
		Members:        nil,
	}

}
