package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	orm "github.com/arcbjorn/arc_bank_system/internal/db/orm"
	"github.com/arcbjorn/arc_bank_system/pkg/utils"
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

func deleteAccountById(t *testing.T, id int64) {
	error := testQueries.DeleteAccount(context.Background(), id)
	require.Nil(t, error)
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
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	arg := orm.UpdateAccountBalanceParams{
		ID:      newAccount.ID,
		Balance: utils.RandomInt(0, 100),
	}

	account, err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Currency, account.Currency)

	require.WithinDuration(t, newAccount.CreatedAt, account.CreatedAt, time.Second)

	require.Equal(t, arg.Balance, account.Balance)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	deleteAccountById(t, newAccount.ID)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := orm.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
