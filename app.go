package main

import (
	"context"
	"database/sql"
	"desktop/internal/appstate"
	catalogApp "desktop/internal/platform/catalog/application"
	catalogInfra "desktop/internal/platform/catalog/infrastructure"
	catalogHandlers "desktop/internal/platform/catalog/infrastructure/handlers"
	iamApp "desktop/internal/platform/iam/application"
	iamPorts "desktop/internal/platform/iam/domain/ports"
	iamInfra "desktop/internal/platform/iam/infrastructure"
	iamHandlers "desktop/internal/platform/iam/infrastructure/handlers"
	inventoryApp "desktop/internal/platform/inventory/application"
	inventoryInfra "desktop/internal/platform/inventory/infrastructure"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
	_ "embed"
	"sync"

	_ "modernc.org/sqlite"
)

//go:embed internal/shared/db/schema.sql
var ddl string

// App struct
type App struct {
	ctx          context.Context
	db           *sql.DB
	queries      *db.Queries
	mu           sync.RWMutex
	appStateRepo iamPorts.AppStateRepository
	sessionState appstate.SessionState
	// Handlers
	CategoryHandler     catalogHandlers.CatalogCategoryHandler
	ProductHandler      catalogHandlers.CatalogProductHandler
	SessionHandler      iamHandlers.IamSessionHandler
	OrganizationHanlder iamHandlers.IamOrganizationHandler
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	conn, err := dbutils.NewDatabase("db/pos.db", ddl, &ctx)
	if err != nil {
		panic(err)
	}

	a.db = conn
	a.queries = db.New(conn)

	// IAM - Repositories & Container
	a.initIamHandlers()

	// Inventory - Repositories & Container
	inventoryRepositories := inventoryInfra.NewRepositoriesContainer(a.queries)

	inventoryUseCasesContainer := inventoryApp.NewUseCasesContainer(
		inventoryRepositories.InventoryItemRepository,
	)

	// Catalog - Repositories & Container
	a.iniCatalogHandlers(inventoryUseCasesContainer)

	a.initAppSession()
}

func (a *App) initAppSession() {
	appState, err := a.appStateRepo.GetWithAggregates(a.ctx)

	if err != nil {
		panic(err)
	}

	sessionState := appstate.BuildSessionState(appState.Operator, appState.Organization)

	a.sessionState = sessionState
}

func (a *App) initIamHandlers() {
	txManager := dbutils.NewSqliteTxManager(a.db, a.queries)

	iamRepositories := iamInfra.NewRepositoryContainer(a.queries)

	iamAppContainer := iamApp.NewUseCasesContainer(
		txManager,
		iamRepositories.OperatorRepository,
		iamRepositories.OrganizationRepository,
		iamRepositories.OrganizationMemberRepository,
		iamRepositories.OrganizationTeamRepository,
	)

	iamQueriesContainer := iamApp.NewQueriesContainer(iamRepositories.OrganizationRepository)

	iamHandlers := iamInfra.NewHandlerContainer(
		iamAppContainer,
		iamQueriesContainer,
		a.onStateChange,
		func() context.Context { return a.ctx },
		func() *appstate.SessionState { return &a.sessionState },
	)

	a.appStateRepo = iamRepositories.AppStateRepository
	a.OrganizationHanlder = iamHandlers.OrganizationHandler
	a.SessionHandler = iamHandlers.SessionHandler
}

func (a *App) iniCatalogHandlers(inventoryUseCasesContainer *inventoryApp.UseCasesContainer) {
	txManager := dbutils.NewSqliteTxManager(a.db, a.queries)

	catalogRepositories := catalogInfra.NewRepositoryContainer(a.queries)

	catalogUseCasesContainer := catalogApp.NewUseCasesContainer(
		txManager,
		catalogRepositories.CategoryRepository,
		catalogRepositories.ProductRepository,
		*inventoryUseCasesContainer.CreateInventoryItem,
		*inventoryUseCasesContainer.UpdateInventoryItem,
	)

	catalogQueriesContainer := catalogApp.NewQueriesContainer(a.queries)

	catalogHandlers := catalogInfra.NewHandlerContainer(
		catalogUseCasesContainer,
		catalogQueriesContainer,
		func() context.Context { return a.ctx },
		func() *appstate.SessionState { return &a.sessionState },
	)

	a.CategoryHandler = catalogHandlers.CategoryHandler
	a.ProductHandler = catalogHandlers.ProductHandler
}

func (a *App) GetAppSession() appstate.SessionStateResponse {
	return appstate.SessionStateResponseFromDomain(&a.sessionState)
}

func (a *App) onStateChange(e appstate.StateChangeEvent) {
	a.mu.Lock()
	defer a.mu.Unlock()

	switch e.Type {
	case appstate.SignUp:
		a.sessionState.SetActiveOperator(e.Operator)
		a.appStateRepo.Update(a.ctx, a.sessionState.ToAppState())

	case appstate.SignIn:
		a.sessionState.SetActiveOperator(e.Operator)
		a.appStateRepo.Update(a.ctx, a.sessionState.ToAppState())

	case appstate.SignOut:
		a.sessionState.Clear()
		a.appStateRepo.Update(a.ctx, a.sessionState.ToAppState())

	case appstate.SwitchOrganization:
		a.sessionState.SetActiveOrganization(e.Organization)
		a.appStateRepo.Update(a.ctx, a.sessionState.ToAppState())
	}

}
