package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang-projects/simplebank/util"
	"github.com/stretchr/testify/require"
)


func createRandomTransfer(t *testing.T) Transfer{
	accounts := getListOfAllAccounts(t)
	n := len(accounts)
	account1 := accounts[rand.Intn(n)]
	account2 := accounts[rand.Intn(n)]
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T){
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(),transfer1.ID)
	require.NoError(t, err)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}


func TestListTransfers(t *testing.T){
	for i:= 0;i< 5;i++ {
		createRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(),arg)
	require.NoError(t, err)
	require.Len(t,transfers,5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}