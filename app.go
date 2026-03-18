package main

import (
	"context"
	"desktop/internal/iam"
	iamAggregates "desktop/internal/iam/domain/aggregates"
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
	SessionState *iamAggregates.AppState
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

	a.IamHandler = *iam.NewIamHandler(ctx, queries, func(appState *iamAggregates.AppState) {
		a.mu.Lock()
		defer a.mu.Unlock()

		if appState == nil {
			log.Println("✗ error getting app state")
			return
		}

		a.SessionState = appState
	})

	log.Println("✓ IamHandler creado")

	// if err != nil {
	// 	log.Println("✗ error appState:", err)
	// 	return
	// }

	_, errAppState := a.IamHandler.GetAppState()

	if errAppState != nil {
		log.Println("✗ error getting app state:", err)
		return
	}

}
