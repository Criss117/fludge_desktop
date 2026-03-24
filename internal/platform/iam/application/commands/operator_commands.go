package commands

type SignIn struct {
	Username string `json:"username"`
	PIN      string `json:"pin"`
}

type RegisterRootOperator struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	PIN      string `json:"pin"`
}

type SwitchOrg struct {
	OrganizationID string `json:"organizationId"`
}
