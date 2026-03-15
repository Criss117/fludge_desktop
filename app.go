package main

import (
	"context"
	"desktop/internal/appstate"
	"desktop/internal/iam"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
	_ "embed"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

//go:embed internal/shared/db/schema.sql
var ddl string

// App struct
type App struct {
	ctx          context.Context
	IamHandler   iam.IamHandler
	SessionState *appstate.SessionState
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
	log.Println("▶ abriendo DB...")

	conn, err := platform.NewDatabase("pos.db", ddl, &ctx)
	if err != nil {
		log.Println("✗ error DB:", err)
		return
	}
	log.Println("✓ DB abierta")

	queries := db.New(conn)

	a.IamHandler = *iam.NewIamHandler(ctx, queries, func(event iam.StateChangeEvent) {
		a.mu.Lock()
		defer a.mu.Unlock()

		switch event.Type {
		case iam.OnStateChangeTypeSignUp:
			a.SessionState = appstate.BuildSessionState(
				nil,
				event.Operator,
				nil,
				nil,
			)
		case iam.OnStateChangeTypeSignIn:
			a.SessionState = appstate.BuildSessionState(
				nil,
				event.Operator,
				nil,
				event.Teams,
			)
		}
	})

	log.Println("✓ authHandler creado")

	// if err != nil {
	// 	log.Println("✗ error appState:", err)
	// 	return
	// }

}

// Greet returns a greeting for the given name
func (a *App) GetSession() *appstate.ResponseSessionState {
	return appstate.ResponseSessionStateFromDomain(
		a.SessionState.ActiveOrganization,
		a.SessionState.ActiveOperator.Operator,
	)
}
