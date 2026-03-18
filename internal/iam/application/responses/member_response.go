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
	primitiveMember := member.ToValues()

	return &MemberResponse{
		ID:             primitiveMember.ID,
		OrganizationID: primitiveMember.OrganizationID,
		OperatorID:     primitiveMember.OperatorID,
		Role:           primitiveMember.Role,
		CreatedAt:      platform.TimeToInt64(primitiveMember.CreatedAt),
		UpdatedAt:      platform.TimeToInt64(primitiveMember.UpdatedAt),
		DeletedAt:      platform.TimeToInt64Nullable(primitiveMember.DeletedAt),
	}
}
