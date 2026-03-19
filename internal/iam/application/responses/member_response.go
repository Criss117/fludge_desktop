package responses

import (
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type MemberResponse struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	OperatorID     string `json:"operatorId"`
	Role           string `json:"role"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	DeletedAt      *int64 `json:"deletedAt"`
}

func MemberResponseFromDomain(member *aggregates.Member) *MemberResponse {

	return &MemberResponse{
		ID:             member.ID,
		OrganizationID: member.OrganizationID,
		OperatorID:     member.OperatorID,
		Role:           member.Role.Value(),
		CreatedAt:      platform.TimeToInt64(member.CreatedAt),
		UpdatedAt:      platform.TimeToInt64(member.UpdatedAt),
		DeletedAt:      platform.TimeToInt64Nullable(member.DeletedAt),
	}
}
