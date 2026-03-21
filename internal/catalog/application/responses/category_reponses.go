package responses

import (
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type CategoryResponse struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	OrganizationID string  `json:"organizationId"`
	CreatedAt      int64   `json:"createdAt"`
	UpdatedAt      int64   `json:"updatedAt"`
	DeletedAt      *int64  `json:"deletedAt"`
}

func CategoryResponseFromDomain(category *aggregates.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:             category.ID,
		Name:           category.Name,
		Description:    category.Description,
		OrganizationID: category.OrganizationID,
		CreatedAt:      platform.ToMillis(category.CreatedAt),
		UpdatedAt:      platform.ToMillis(category.UpdatedAt),
		DeletedAt:      platform.TimeToInt64Nullable(category.DeletedAt),
	}
}
