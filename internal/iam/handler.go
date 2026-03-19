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
	"log"
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

type OnStateChange func(appState *aggregates.AppState)

type IamHandler struct {
	ctx                    context.Context
	onStateChange          OnStateChange
	organizationRepository ports.OrganizationRepository
	operatorRepository     ports.OperatorRepository
	appStateRepository     ports.AppStateRepository
	memberRepository       ports.MemberRepository
	getAppStateUseCase     *usecases.GetAppStateUseCase
}

func NewIamHandler(
	ctx context.Context,
	queries *db.Queries,
	onStateChange OnStateChange,
) *IamHandler {
	organizationRepository := repositories.NewSqliteOrganizationRepository(queries)
	operatorRepository := repositories.NewSqliteOperatorRepository(queries, organizationRepository)
	memberRepository := repositories.NewSqliteMemberRepository(queries)
	appStateRepository := repositories.NewSqliteAppStateRepository(
		queries,
		organizationRepository,
		operatorRepository,
	)

	getAppStateUseCase := usecases.NewGetAppStateUseCase(appStateRepository, organizationRepository)

	return &IamHandler{
		ctx:                    ctx,
		onStateChange:          onStateChange,
		organizationRepository: organizationRepository,
		operatorRepository:     operatorRepository,
		appStateRepository:     appStateRepository,
		memberRepository:       memberRepository,
		getAppStateUseCase:     getAppStateUseCase,
	}
}

func (h *IamHandler) SignUp(signUpDtp *commands.SignUpCommand) (*responses.ResponseAppState, error) {
	signUpUseCase := usecases.NewSignUpUseCase(h.operatorRepository, h.appStateRepository)

	newAppState, err := signUpUseCase.Execute(h.ctx, signUpDtp)

	if err != nil {
		return nil, err
	}

	h.onStateChange(newAppState)

	return responses.ResponseAppStateFromDomain(newAppState), nil
}

func (h *IamHandler) SignIn(signInDto *commands.SignInCommand) (*responses.ResponseAppState, error) {
	log.Println("✓ SignIn 1")
	signInUseCase := usecases.NewSignInUseCase(h.operatorRepository, h.appStateRepository)

	log.Println("✓ SignIn 2")
	newAppState, err := signInUseCase.Execute(h.ctx, signInDto)

	if err != nil {
		return nil, err
	}

	log.Println("✓ SignIn 3")
	h.onStateChange(newAppState)

	log.Println("✓ SignIn 4")
	return responses.ResponseAppStateFromDomain(newAppState), nil
}

func (h *IamHandler) SignOut() error {
	signOutUseCase := usecases.NewSignOutUseCase(h.appStateRepository)
	err := signOutUseCase.Execute(h.ctx)

	if err != nil {
		return err
	}

	appState, err := h.getAppStateUseCase.Execute(h.ctx)

	if err != nil {
		return err
	}

	h.onStateChange(appState)

	return nil
}

func (h *IamHandler) SwitchOrganization(
	switchOrganizationDto *commands.SwitchOrganizationCommand,
) (*responses.ResponseAppState, error) {
	switchOrganizationUseCase := usecases.NewSwitchOrganizationUseCase(h.organizationRepository, h.appStateRepository)

	newAppState, err := switchOrganizationUseCase.Execute(h.ctx, switchOrganizationDto)

	if err != nil {
		return nil, err
	}

	h.onStateChange(newAppState)

	return responses.ResponseAppStateFromDomain(newAppState), nil
}

func (h *IamHandler) GetAppState() (*responses.ResponseAppState, error) {
	appState, err := h.getAppStateUseCase.Execute(h.ctx)

	if err != nil {
		return nil, err
	}

	if appState == nil {
		return nil, nil
	}

	h.onStateChange(appState)

	return responses.ResponseAppStateFromDomain(appState), nil
}

func (h *IamHandler) FindManyOrganizationsByOperatorId(operatorId string) ([]*responses.OrganizationResponse, error) {
	organizations, err := h.organizationRepository.FindByOperator(h.ctx, operatorId)

	if err != nil {
		return nil, err
	}

	responseOrganizations := make([]*responses.OrganizationResponse, len(organizations))

	for index, organization := range organizations {
		responseOrganizations[index] = responses.OrganizationResponseFromDomain(organization)
	}

	return responseOrganizations, nil
}

func (h *IamHandler) RegisterOrganization(
	registerOrganizationDto *commands.RegisterOrganizationCommand,
) (*responses.OrganizationResponse, error) {
	registerOrganizationUseCase := usecases.NewRegisterOrganizationUseCase(
		h.organizationRepository,
		h.memberRepository,
	)

	appState, err := h.getAppStateUseCase.Execute(h.ctx)

	newOrganization, err := registerOrganizationUseCase.Execute(h.ctx,
		*appState.ActiveOperator,
		registerOrganizationDto,
	)

	if err != nil {
		return nil, err
	}

	_, errSwitchOrganization := h.SwitchOrganization(&commands.SwitchOrganizationCommand{
		OrganizationID: newOrganization.ID,
	})

	if errSwitchOrganization != nil {
		return nil, err
	}

	h.onStateChange(nil)

	return responses.OrganizationResponseFromDomain(newOrganization), nil
}
