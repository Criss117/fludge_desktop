package infrastructure

import (
	"desktop/internal/platform/iam/application"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/platform/iam/infrastructure/handlers"
	"desktop/internal/platform/iam/infrastructure/repositories"
	"desktop/internal/shared/db"
	"desktop/internal/shared/infrastructure"
)

type RepositoryContainer struct {
	AppStateRepository           ports.AppStateRepository
	OperatorRepository           ports.OperatorRepository
	OrganizationRepository       ports.OrganizationRepository
	OrganizationMemberRepository ports.OrganizationMemberRepository
	OrganizationTeamRepository   ports.OrganizationTeamRepository
}

type HandlerContainer struct {
	OrganizationHandler handlers.IamOrganizationHandler
	SessionHandler      handlers.IamSessionHandler
}

func NewRepositoryContainer(querires *db.Queries) *RepositoryContainer {
	operatorRepository := repositories.NewSqliteOperatorRepository(querires)
	memberRepository := repositories.NewSqliteOrganizationMemberRepository(querires)
	teamRepository := repositories.NewSqliteOrganizationTeamRepository(querires)
	organizationRepository := repositories.NewSqliteOrganizationRepository(querires, memberRepository, teamRepository)
	appStateRepository := repositories.NewSqliteAppRepository(
		querires,
		organizationRepository,
		operatorRepository,
	)

	return &RepositoryContainer{
		AppStateRepository:           appStateRepository,
		OperatorRepository:           operatorRepository,
		OrganizationRepository:       organizationRepository,
		OrganizationMemberRepository: memberRepository,
		OrganizationTeamRepository:   teamRepository,
	}
}

func NewHandlerContainer(
	app *application.UseCasesContainer,
	queries *application.QueriesContainer,
	onStateChange handlers.OnStateChange,
	getCtx infrastructure.GetCtxFunc,
	getSession infrastructure.GetSessionFunc,
) *HandlerContainer {
	organizationHanlder := handlers.NewIamOrganizationHandler(
		app,
		queries,
		onStateChange,
		getCtx,
		getSession,
	)

	sessionHandler := handlers.NewIamSessionHandler(
		app,
		queries,
		onStateChange,
		getCtx,
		getSession,
	)

	return &HandlerContainer{
		OrganizationHandler: organizationHanlder,
		SessionHandler:      sessionHandler,
	}
}
