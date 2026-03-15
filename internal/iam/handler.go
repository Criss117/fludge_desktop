package iam

import (
	"context"
	"desktop/internal/iam/application/commands"
	"desktop/internal/iam/application/responses"
	"desktop/internal/iam/application/usecases"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/ports"
	"desktop/internal/iam/infrastructure/repositories"
	"desktop/internal/shared/db"
)

type OnStateChangeType string
type StateChangeEvent struct {
	Type         OnStateChangeType
	Operator     *aggregates.Operator
	Organization *aggregates.Organization
	Member       *aggregates.Member
	Teams        []*aggregates.Team
}

const (
	OnStateChangeTypeSwitchOrganization OnStateChangeType = "switch_organization"
	OnStateChangeTypeSignUp             OnStateChangeType = "sign_up"
	OnStateChangeTypeSignIn             OnStateChangeType = "sign_in"
	OnStateChangeTypeSignOut            OnStateChangeType = "sign_out"
	OnStateChangeTypeRefetch            OnStateChangeType = "refetch"
)

type OnStateChange func(event StateChangeEvent)

type IamHandler struct {
	ctx                    context.Context
	onStateChange          OnStateChange
	organizationRepository ports.OrganizationRepository
	operatorRepository     ports.OperatorRepository
	appStateRepository     ports.AppStateRepository
}

func NewIamHandler(ctx context.Context, queries *db.Queries, onStateChange OnStateChange) *IamHandler {
	organizationRepository := repositories.NewSqliteOrganizationRepository(queries)
	operatorRepository := repositories.NewSqliteOperatorRepository(queries)
	appStateRepository := repositories.NewSqliteAppStateRepository(
		queries,
		organizationRepository,
		operatorRepository,
	)

	return &IamHandler{
		ctx:                    ctx,
		onStateChange:          onStateChange,
		organizationRepository: organizationRepository,
		operatorRepository:     operatorRepository,
		appStateRepository:     appStateRepository,
	}
}

func (h *IamHandler) SignUp(signUpDtp *commands.SignUpCommand) (*responses.OperatorResponse, error) {
	signUpUseCase := usecases.NewSignUpUseCase(h.operatorRepository, h.appStateRepository)

	operator, err := signUpUseCase.Execute(h.ctx, signUpDtp)

	if err != nil {
		return nil, err
	}

	h.onStateChange(StateChangeEvent{
		Type:     OnStateChangeTypeSignUp,
		Operator: operator,
	})

	return responses.OperatorResponseFromDomain(operator), nil
}

func (h *IamHandler) SignIn(signInDto *commands.SignInCommand) (*responses.SignInResponse, error) {
	signInUseCase := usecases.NewSignInUseCase(h.operatorRepository, h.appStateRepository)

	signInResponse, err := signInUseCase.Execute(h.ctx, signInDto)

	if err != nil {
		return nil, err
	}

	h.onStateChange(StateChangeEvent{
		Type:     OnStateChangeTypeSignIn,
		Operator: signInResponse.Operator,
		Teams:    signInResponse.Teams,
	})

	operatorResponse := responses.OperatorResponseFromDomain(signInResponse.Operator)
	teamResponse := make([]responses.TeamResponse, len(signInResponse.Teams))

	for i, team := range signInResponse.Teams {
		teamResponse[i] = *responses.TeamResponseFromDomain(team)
	}

	h.onStateChange(StateChangeEvent{
		Type:     OnStateChangeTypeSignIn,
		Operator: signInResponse.Operator,
		Teams:    signInResponse.Teams,
	})

	return &responses.SignInResponse{
		ActiveOperator: *operatorResponse,
		ActiveTeams:    teamResponse,
	}, nil
}
