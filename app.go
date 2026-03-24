package main

import (
	"context"
	"database/sql"
	"desktop/internal/appstate"
	"desktop/internal/platform/iam"
	iamApp "desktop/internal/platform/iam/application"
	iamPorts "desktop/internal/platform/iam/domain/ports"
	iamRepos "desktop/internal/platform/iam/infrastructure/repositories"

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
	SessionState *appstate.SessionState
	IamHandler   iam.IamHandler
	AppStateRepo iamPorts.AppStateRepository
	mu           sync.RWMutex
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

	// IAM - Repositories
	operatorRepo := iamRepos.NewSqliteOperatorRepository(a.queries)
	appStateRepo := iamRepos.NewSqliteAppRepository(a.queries)
	organizationRepo := iamRepos.NewSqliteOrganizationRepository(a.queries)
	memberRepo := iamRepos.NewSqliteOrganizationMemberRepository(a.queries)
	teamRepo := iamRepos.NewSqliteOrganizationTeamRepository(a.queries)

	iamContainer := iamApp.NewContainer(
		txManager,
		operatorRepo,
		organizationRepo,
		memberRepo,
		teamRepo,
	)

	a.AppStateRepo = appStateRepo

	// IAM - Handler
	a.IamHandler = *iam.NewIamHandler(
		func() context.Context { return a.ctx },
		a.onStateChange,
		iamContainer,
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
