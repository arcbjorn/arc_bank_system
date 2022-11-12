package db

import (
	"context"
	"database/sql"
	"fmt"
)

// functions to execute queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// executes function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rollbackEror := tx.Rollback(); rollbackEror != nil {
			return fmt.Errorf("tx error: %v, rollbsack error: %v", err, rollbackEror)
		}

		return err
	}

	return tx.Commit()
}
