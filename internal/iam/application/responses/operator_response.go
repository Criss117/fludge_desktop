package responses

import (
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type OperatorOrganizationResponse struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type OperatorResponse struct {
	ID         string                          `json:"id"`
	Name       string                          `json:"name"`
	Email      string                          `json:"email"`
	Username   string                          `json:"username"`
	Pin        string                          `json:"pin"`
	IsRoot     bool                            `json:"isRoot"`
	CreatedAt  int64                           `json:"createdAt"`
	UpdatedAt  int64                           `json:"updatedAt"`
	DeletedAt  *int64                          `json:"deletedAt"`
	IsMemberIn []*OperatorOrganizationResponse `json:"isMemberIn"`
}

func OperatorResponseFromDomain(operator *aggregates.Operator) *OperatorResponse {

	organizations := make([]*OperatorOrganizationResponse, len(operator.IsMemberIn))

	for i, organization := range operator.IsMemberIn {
		organizations[i] = &OperatorOrganizationResponse{
			ID:   organization.ID,
			Slug: organization.Slug,
			Name: organization.Name,
		}
	}

	return &OperatorResponse{
		ID:         operator.ID,
		Name:       operator.Name,
		Email:      operator.Email.Value(),
		Username:   operator.Username,
		Pin:        operator.Pin.Value(),
		IsRoot:     operator.Root,
		CreatedAt:  platform.TimeToInt64(operator.CreatedAt),
		UpdatedAt:  platform.TimeToInt64(operator.UpdatedAt),
		DeletedAt:  platform.TimeToInt64Nullable(operator.DeletedAt),
		IsMemberIn: organizations,
	}
}
