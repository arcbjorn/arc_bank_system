package db

import (
	"context"
	"testing"

	orm "github.com/arcbjorn/arc_bank_system/internal/db/orm"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := orm.CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
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
}