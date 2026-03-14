package auth

import (
	"context"
	"desktop/internal/auth/application"
	"desktop/internal/auth/application/dtos"
	"desktop/internal/auth/domain"
	"desktop/internal/auth/infrastruture"
	"desktop/internal/shared/db"
)

type OnStateChange func(*domain.AppState)

type AuthHandler struct {
	OperatorRepository      domain.OperatorRepository
	OnStateChange           OnStateChange
	FindAllOperatorsUseCase *application.FindAllOperatorsUseCase
	SignUpUseCase           *application.SignUpUseCase
	FindAppStateUseCase     *application.FindAppStateUseCase
	SignInUseCase           *application.SignInUseCase
}

func NewAuthHandler(ctx *context.Context, queries *db.Queries, callback OnStateChange) *AuthHandler {
	operatorRepository := infrastruture.NewSQLiteOperatorRepository(ctx, queries)
	appStateRepository := infrastruture.NewSQLiteAppStateRepository(ctx, queries)

	// use cases
	findAllOperatorsUseCase := application.NewFindAllOperatorsUseCase(operatorRepository)
	signUpUseCase := application.NewSignUpUseCase(operatorRepository, appStateRepository)
	findAppStateUseCase := application.NewFindAppStateUseCase(appStateRepository)
	signInUseCase := application.NewSignInUseCase(operatorRepository, appStateRepository)

	return &AuthHandler{
		OperatorRepository:      operatorRepository,
		OnStateChange:           callback,
		FindAllOperatorsUseCase: findAllOperatorsUseCase,
		SignUpUseCase:           signUpUseCase,
		FindAppStateUseCase:     findAppStateUseCase,
		SignInUseCase:           signInUseCase,
	}
}

func (ah *AuthHandler) FindAllOperators() ([]*dtos.SummaryOperatorDTO, error) {
	return ah.FindAllOperatorsUseCase.Execute()
}

func (ah *AuthHandler) CreateOperator(operator dtos.SignupDTO) (*dtos.AppStateDTO, error) {
	newAppState, err := ah.SignUpUseCase.Execute(operator)

	if err != nil {
		return nil, err
	}

	ah.OnStateChange(newAppState)

	return dtos.DomainToAppStateDTO(newAppState), nil
}

func (ah *AuthHandler) FindAppState() (*dtos.AppStateDTO, error) {
	return ah.FindAppStateUseCase.Execute()
}

func (ah *AuthHandler) SignIn(signInDto dtos.SignInDTO) (*dtos.AppStateDTO, error) {
	newAppState, err := ah.SignInUseCase.Execute(signInDto.Username, signInDto.Pin)

	if err != nil {
		return nil, err
	}

	ah.OnStateChange(newAppState)

	return dtos.DomainToAppStateDTO(newAppState), nil
}
