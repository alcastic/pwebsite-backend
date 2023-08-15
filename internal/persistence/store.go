package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/alcastic/pwebsite-backend/internal/generated/sqlc"
)

type Store struct {
	*sqlc.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	store := &Store{
		Queries: sqlc.New(db),
		db:      db,
	}
	return store
}

func (s *Store) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(tx)
	txErr := fn(queries)

	if txErr != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %s, rollback error: %s", txErr, rbErr)
		}
		return txErr
	}

	return tx.Commit()
}
