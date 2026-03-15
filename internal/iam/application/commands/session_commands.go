package commands

type SignInCommand struct {
	Username string `json:"username"`
	PIN      string `json:"pin"`
}

type SignUpCommand struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	PIN      string `json:"pin"`
}

type SwitchOrgCommand struct {
	OrganizationID string `json:"organizationId"`
}
