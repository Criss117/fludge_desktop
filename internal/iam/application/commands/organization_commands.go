package commands

type RegisterOrganizationCommand struct {
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	LegalName    string  `json:"legalName"`
	Address      string  `json:"address"`
	Logo         *string `json:"logo"`
	ContactPhone *string `json:"contactPhone"`
	ContactEmail *string `json:"contactEmail"`
}
