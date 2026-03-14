package platform

import (
	"context"
	"database/sql"
)

func NewDatabase(path string, ddl string, ctx *context.Context) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", path+"?_foreign_keys=on&_journal_mode=WAL")

	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(1) // SQLite no soporta writes concurrentes

	if _, err := conn.ExecContext(*ctx, ddl); err != nil {
		return nil, err
	}

	return conn, nil
}
