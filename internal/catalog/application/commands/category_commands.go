package commands

type CreateCategoryCommand struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
