package responses

type SignUpResponse struct {
	ActiveOrganization OrganizationResponse `json:"activeOrganization"`
}

type SignInResponse struct {
	ActiveOperator OperatorResponse `json:"activeOperator"`
	ActiveTeams    []TeamResponse   `json:"activeTeams"`
}
