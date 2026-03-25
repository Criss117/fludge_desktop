package responses

import (
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/shared/db/dbutils"
)

type Category struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	OrganizationID string  `json:"organization_id"`
	CreatedAt      int64   `json:"created_at"`
	UpdatedAt      int64   `json:"updated_at"`
	DeletedAt      *int64  `json:"deleted_at"`
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
