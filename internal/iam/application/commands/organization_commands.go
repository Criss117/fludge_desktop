package commands

type RegisterOrganizationCommand struct {
	Name         string  `json:"name"`
	LegalName    string  `json:"legalName"`
	Address      string  `json:"address"`
	Logo         *string `json:"logo"`
	ContactPhone *string `json:"contactPhone"`
	ContactEmail *string `json:"contactEmail"`
}
