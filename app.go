package main

import (
	"context"
	"desktop/internal/auth"
	authDomain "desktop/internal/auth/domain"
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
	ctx         context.Context
	authHandler auth.AuthHandler
	appState    *authDomain.AppState
	mu          sync.RWMutex
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// // startup is called when the app starts. The context is saved
// // so we can call the runtime methods
// func (a *App) startup(ctx context.Context) {
// 	a.ctx = ctx

// 	conn, err := platform.NewDatabase("pos.db", ddl, &ctx)

// 	if err != nil {
// 		log.Fatal("error abriendo DB:", err)
// 	}

// 	queries := db.New(conn)

// 	a.authHandler = *auth.NewAuthHandler(&ctx, queries, func(as *authDomain.AppState) {
// 		a.mu.Lock()
// 		a.appState = as
// 		a.mu.Unlock()
// 	})

// 	state, err := a.authHandler.FindAppStateUseCase.ExecuteWithoutParsing()
// 	if err != nil {
// 		log.Fatal("error obteniendo estado de la aplicación:", err)
// 	}

// 	a.mu.Lock()
// 	a.appState = state
// 	a.mu.Unlock()
// }

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

	a.authHandler = *auth.NewAuthHandler(&ctx, queries, func(as *authDomain.AppState) {
		a.mu.Lock()
		a.appState = as
		a.mu.Unlock()
	})

	log.Println("✓ authHandler creado")

	state, err := a.authHandler.FindAppStateUseCase.ExecuteWithoutParsing()
	if err != nil {
		log.Println("✗ error appState:", err)
		return
	}
	log.Println("✓ appState:", state)

	a.mu.Lock()
	a.appState = state
	a.mu.Unlock()
}
