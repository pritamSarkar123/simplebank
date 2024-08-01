package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang-projects/simplebank/util"
	"github.com/stretchr/testify/require"
)

func getListOfAllAccounts(t *testing.T) []Account{
	arg := ListAccountsParams{
		Limit: 10,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(),arg)
	require.NoError(t, err)
	require.Len(t,accounts,10)
	return accounts
}
func createRandomEntry(t *testing.T) Entry{
	accounts := getListOfAllAccounts(t)
	n := len(accounts)
	account := accounts[rand.Intn(n)]
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

func TestCreateEntry(t *testing.T){
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T){
	enrty1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(),enrty1.ID)
	require.NoError(t, err)
	require.Equal(t, enrty1.ID, entry2.ID)
	require.Equal(t, enrty1.AccountID, entry2.AccountID)
	require.Equal(t, enrty1.Amount, entry2.Amount)
	require.WithinDuration(t, enrty1.CreatedAt, enrty1.CreatedAt, time.Second)
}


func TestListEntries(t *testing.T){
	for i:= 0;i< 5;i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(),arg)
	require.NoError(t, err)
	require.Len(t,entries,5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}