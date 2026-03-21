package responses

import (
	"desktop/internal/catalog/domain/aggregates"
)

type CategoryResponse struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	OrganizationID string  `json:"organizationId"`
}

func CategoryResponseFromDomain(category *aggregates.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:             category.ID,
		Name:           category.Name,
		Description:    category.Description,
		OrganizationID: category.OrganizationID,
	}
}
