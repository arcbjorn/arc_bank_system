package db

import (
	"context"
	"testing"

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

	error := testQueries.DeleteAccount(context.Background(), account.ID)
	require.Nil(t, error)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
