package responses

import (
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/shared/db/dbutils"
)

type Category struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	OrganizationID string  `json:"organizationId"`
	CreatedAt      int64   `json:"createdAt"`
	UpdatedAt      int64   `json:"updatedAt"`
	DeletedAt      *int64  `json:"deletedAt"`
}

func CategoryResponseFromDomain(category *aggregates.Category) *Category {
	return &Category{
		ID:             category.ID,
		Name:           category.Name,
		Description:    category.Description,
		OrganizationID: category.OrganizationID,
		CreatedAt:      dbutils.TimeToInt64(category.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(category.UpdatedAt),
		DeletedAt:      dbutils.TimeToInt64Nullable(category.DeletedAt),
	}
}
