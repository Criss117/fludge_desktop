package domain

type OperatorRepository interface {
	FinAll() ([]*Operator, error)
	FindOneByUsername(username string) (*Operator, error)
	FindOneByEmail(email string) (*Operator, error)
	FindManyByUsernameOrEmail(username, email string) ([]*Operator, error)

	Create(operator Operator) error
}

type AppStateRepository interface {
	FindAppState() (*AppState, error)
	Update(appState AppState) error
}
