package responses

import (
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/shared/db/dbutils"
)

type Organization struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Logo         *string `json:"logo"`
	Metadata     []byte  `json:"metadata"`
	LegalName    string  `json:"legalName"`
	Address      string  `json:"address"`
	ContactPhone *string `json:"contactPhone"`
	ContactEmail *string `json:"contactEmail"`
	CreatedAt    int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAt"`
	DeletedAt    *int64  `json:"deletedAt"`
	Members      []*Member
	Teams        []*Team
}

func OrganizationResponseFromDomain(organization *aggregates.Organization) *Organization {

	contactEmail := organization.ContactEmail.Value()

	return &Organization{
		ID:           organization.ID,
		Name:         organization.Name,
		Slug:         organization.Slug.Value(),
		Logo:         organization.Logo,
		LegalName:    organization.LegalName,
		Address:      organization.Address,
		ContactPhone: organization.ContactPhone,
		ContactEmail: &contactEmail,
		CreatedAt:    dbutils.TimeToInt64(organization.CreatedAt),
		UpdatedAt:    dbutils.TimeToInt64(organization.UpdatedAt),
		DeletedAt:    dbutils.TimeToInt64Nullable(organization.DeletedAt),
	}
}
