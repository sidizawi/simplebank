package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sidizawi/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 7; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  4,
		Offset: 1,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, int(arg.Limit))

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}
