package commands

type RegisterOrganization struct {
	Name         string  `json:"name"`
	LegalName    string  `json:"legalName"`
	Address      string  `json:"address"`
	Logo         *string `json:"logo"`
	ContactPhone *string `json:"contactPhone"`
	ContactEmail *string `json:"contactEmail"`
}

type UpdateOrganization struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	LegalName    string  `json:"legalName"`
	Address      string  `json:"address"`
	Logo         *string `json:"logo"`
	ContactPhone *string `json:"contactPhone"`
	ContactEmail *string `json:"contactEmail"`
}

type SwitchOrganization struct {
	OrganizationId string `json:"organizationId"`
}

type FindOneOrganization struct {
	OrganizationId string `json:"organizationId"`
}
