package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	orm "github.com/arcbjorn/arc_bank_system/internal/db/orm"

	"github.com/arcbjorn/arc_bank_system/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account orm.Account) orm.Entry {
	arg := orm.CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomInt(0, 100),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func deleteEntryById(t *testing.T, id int64) {
	error := testQueries.DeleteEntry(context.Background(), id)
	require.Nil(t, error)
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	newAccount := createRandomAccount(t)
	newEntry := createRandomEntry(t, newAccount)
	entry, err := testQueries.GetEntry(context.Background(), newEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, newEntry.ID, entry.ID)
	require.Equal(t, newEntry.AccountID, entry.AccountID)
	require.Equal(t, newEntry.Amount, entry.Amount)

	require.WithinDuration(t, newEntry.CreatedAt, entry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	newAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, newAccount)
	}

	arg := orm.ListEntriesParams{
		AccountID: newAccount.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}

func TestDeleteEntry(t *testing.T) {
	newAccount := createRandomAccount(t)
	newEntry := createRandomEntry(t, newAccount)

	deleteEntryById(t, newEntry.ID)

	entry, err := testQueries.GetEntry(context.Background(), newEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry)
}
