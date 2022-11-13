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

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
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

// create transfer, create account entries & update account balance
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return nil
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})

		if err != nil {
			return nil
		}

		// update balances

		account1, err := q.GetAccountForUpdate(context.Background(), arg.FromAccountID)
		if err != nil {
			return nil
		}

		result.FromAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount,
		})

		if err != nil {
			return nil
		}

		account2, err := q.GetAccountForUpdate(context.Background(), arg.ToAccountId)
		if err != nil {
			return nil
		}

		result.ToAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
			ID:      arg.ToAccountId,
			Balance: account2.Balance + arg.Amount,
		})

		if err != nil {
			return nil
		}

		return nil
	})

	return result, err
}
