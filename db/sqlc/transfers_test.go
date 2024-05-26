package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	account1, err = testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err = testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.NotEmpty(t, transfer)
	require.NotEmpty(t, account1)
	require.NotEmpty(t, account2)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
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
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, 0)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	args := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: transfer1.Amount + 10,
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.Amount, transfer2.Amount-10)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
