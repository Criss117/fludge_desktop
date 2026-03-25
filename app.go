package main

import (
	"context"
	"database/sql"
	"desktop/internal/appstate"
	"desktop/internal/platform/catalog"
	catalogApp "desktop/internal/platform/catalog/application"
	catalogInfra "desktop/internal/platform/catalog/infrastructure"
	"desktop/internal/platform/iam"
	iamApp "desktop/internal/platform/iam/application"
	iamPorts "desktop/internal/platform/iam/domain/ports"
	iamInfra "desktop/internal/platform/iam/infrastructure"
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
	ctx            context.Context
	db             *sql.DB
	queries        *db.Queries
	SessionState   *appstate.SessionState
	IamHandler     iam.IamHandler
	CatalogHandler catalog.CatalogHandler
	AppStateRepo   iamPorts.AppStateRepository
	mu             sync.RWMutex
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	conn, err := dbutils.NewDatabase("pos.db", ddl, &ctx)
	if err != nil {
		panic(err)
	}

	a.db = conn
	a.queries = db.New(conn)

	// TxManager
	txManager := dbutils.NewSqliteTxManager(conn, a.queries)

	// IAM - Repositories & Container
	aimRepositories := iamInfra.NewContainer(a.queries)

	iamAppContainer := iamApp.NewContainer(
		txManager,
		aimRepositories.OperatorRepository,
		aimRepositories.OrganizationRepository,
		aimRepositories.OrganizationMemberRepository,
		aimRepositories.OrganizationTeamRepository,
	)

	a.AppStateRepo = aimRepositories.AppStateRepository

	a.IamHandler = *iam.NewIamHandler(
		iamAppContainer,
		a.onStateChange,
		func() context.Context { return a.ctx },
		func() *appstate.SessionState { return a.SessionState },
	)

	// Inventory - Repositories & Container
	inventoryRepositories := inventoryInfra.NewRepositoriesContainer(a.queries)

	inventoryUseCasesContainer := inventoryApp.NewUseCasesContainer(
		inventoryRepositories.InventoryItemRepository,
	)

	// Catalog - Repositories & Container
	catalogRepositories := catalogInfra.NewContainer(a.queries)

	catalogUseCasesContainer := catalogApp.NewUseCasesContainer(
		txManager,
		catalogRepositories.CategoryRepository,
		catalogRepositories.ProductRepository,
		*inventoryUseCasesContainer.CreateInventoryItem,
		*inventoryUseCasesContainer.UpdateInventoryItem,
	)

	catalogQueriesContainer := catalogApp.NewQueriesContainer(a.queries)

	a.CatalogHandler = *catalog.NewCatalogHandler(
		catalogUseCasesContainer,
		catalogQueriesContainer,
		func() context.Context { return a.ctx },
		func() *appstate.SessionState { return a.SessionState },
	)
}

func (a *App) GetAppSession() *appstate.SessionStateResponse {
	return appstate.SessionStateResponseFromDomain(a.SessionState)
}

func (a *App) onStateChange(e appstate.StateChangeEvent) {
	a.mu.Lock()
	defer a.mu.Unlock()

	switch e.Type {
	case appstate.SignUp:
		a.SessionState.SetActiveOperator(e.Operator)
		a.AppStateRepo.Update(a.ctx, a.SessionState.ToAppState())

	case appstate.SignIn:
		a.SessionState.SetActiveOperator(e.Operator)
		a.AppStateRepo.Update(a.ctx, a.SessionState.ToAppState())

	case appstate.SignOut:
		a.SessionState.Clear()
		a.AppStateRepo.Update(a.ctx, a.SessionState.ToAppState())

	case appstate.SwitchOrganization:
		a.SessionState.SetActiveOrganization(e.Organization)
		a.AppStateRepo.Update(a.ctx, a.SessionState.ToAppState())
	}
}
