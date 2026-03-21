// catalog/domain/aggregates/category.go
package aggregates

import (
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/shared/lib"
	"strings"
	"time"
)

type Category struct {
	ID             string
	Name           string
	Description    *string
	OrganizationID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewCategory(name string, description *string, organizationID string) (*Category, error) {
	if len(strings.TrimSpace(name)) < 2 {
		return nil, derrors.ErrCategoryNameTooShort
	}
	return &Category{
		ID:             lib.GenerateUUID(),
		Name:           strings.TrimSpace(name),
		Description:    description,
		OrganizationID: organizationID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func ReconstituteCategory(
	id, name string,
	description *string,
	organizationID string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) *Category {
	return &Category{
		ID:             id,
		Name:           name,
		Description:    description,
		OrganizationID: organizationID,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}

func (c *Category) UpdateDetails(name string, description *string) error {
	if len(strings.TrimSpace(name)) < 2 {
		return derrors.ErrCategoryNameTooShort
	}

	c.Description = description
	c.Name = strings.TrimSpace(name)
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Category) Delete() {
	now := time.Now()
	c.DeletedAt = &now
	c.UpdatedAt = now
}

func (c *Category) IsActive() bool              { return c.DeletedAt == nil }
func (c *Category) Equals(other *Category) bool { return c.ID == other.ID }
