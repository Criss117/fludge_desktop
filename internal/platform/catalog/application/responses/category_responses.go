package responses

import "desktop/internal/platform/catalog/domain/aggregates"

type Category struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	OrganizationID string  `json:"organization_id"`
}

func CategoryResponseFromDomain(category *aggregates.Category) *Category {
	return &Category{
		ID:             category.ID,
		Name:           category.Name,
		Description:    category.Description,
		OrganizationID: category.OrganizationID,
	}
}
