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
	newAccount := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)

	require.WithinDuration(t, newAccount.CreatedAt, account.CreatedAt, time.Second)

	error := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.Nil(t, error)
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	arg := orm.UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: utils.RandomInt(0, 100),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Currency, account.Currency)

	require.WithinDuration(t, newAccount.CreatedAt, account.CreatedAt, time.Second)

	require.Equal(t, arg.Balance, account.Balance)

	error := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.Nil(t, error)
}
