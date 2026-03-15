package responses

import (
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type OrganizationResponse struct {
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
}

func OrganizationResponseFromDomain(organization *aggregates.Organization) *OrganizationResponse {
	primitiveOrganization := organization.ToValues()

	return &OrganizationResponse{
		ID:           primitiveOrganization.ID,
		Name:         primitiveOrganization.Name,
		Slug:         primitiveOrganization.Slug,
		Logo:         primitiveOrganization.Logo,
		LegalName:    primitiveOrganization.LegalName,
		Address:      primitiveOrganization.Address,
		ContactPhone: primitiveOrganization.ContactPhone,
		ContactEmail: primitiveOrganization.ContactEmail,
		CreatedAt:    platform.TimeToInt64(primitiveOrganization.CreatedAt),
		UpdatedAt:    platform.TimeToInt64(primitiveOrganization.UpdatedAt),
		DeletedAt:    platform.TimeToInt64Nullable(primitiveOrganization.DeletedAt),
	}
}
