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
	Members      []*MemberResponse
	Teams        []*TeamResponse
}

func OrganizationResponseFromDomain(organization *aggregates.Organization) *OrganizationResponse {

	contactEmail := organization.ContactEmail.Value()

	return &OrganizationResponse{
		ID:           organization.ID,
		Name:         organization.Name,
		Slug:         organization.Slug.Value(),
		Logo:         organization.Logo,
		LegalName:    organization.LegalName,
		Address:      organization.Address,
		ContactPhone: organization.ContactPhone,
		ContactEmail: &contactEmail,
		CreatedAt:    platform.TimeToInt64(organization.CreatedAt),
		UpdatedAt:    platform.TimeToInt64(organization.UpdatedAt),
		DeletedAt:    platform.TimeToInt64Nullable(organization.DeletedAt),
	}
}
