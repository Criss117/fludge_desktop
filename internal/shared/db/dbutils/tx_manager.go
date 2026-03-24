package dbutils

import (
	"context"
	"database/sql"
	"desktop/internal/shared/db"
)

type TxManager interface {
	WithTx(ctx context.Context, fn func(q *db.Queries) error) error
}

type SqliteTxManager struct {
	db      *sql.DB
	queries *db.Queries
}

func NewSqliteTxManager(db *sql.DB, queries *db.Queries) *SqliteTxManager {
	return &SqliteTxManager{
		db:      db,
		queries: queries,
	}
}

func (m *SqliteTxManager) WithTx(
	ctx context.Context,
	fn func(q *db.Queries) error,
) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := m.queries.WithTx(tx)

	// Manejo seguro de panic + rollback
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r) // re-lanza el panic
		}
	}()

	if err := fn(qtx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
