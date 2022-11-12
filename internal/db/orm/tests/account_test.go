package db

import (
	"context"
	"testing"
	"time"

	orm "github.com/arcbjorn/arc_bank_system/internal/db/orm"
	"github.com/arcbjorn/arc_bank_system/pkg/utils"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) orm.Account {
	arg := orm.CreateAccountParams{
		Owner:    utils.RandomName(),
		Balance:  utils.RandomInt(0, 100),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	error := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.Nil(t, error)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := orm.UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomInt(0, 100),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	require.Equal(t, arg.Balance, account2.Balance)

	error := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.Nil(t, error)
}
