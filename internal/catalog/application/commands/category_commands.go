package commands

type CreateCategory struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateCategory struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type DeleteManyCategories struct {
	IDs []string `json:"ids"`
}
