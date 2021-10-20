package transaction

import (
	"github.com/ozonmp/omp-bot/internal/model/bank/transaction"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var mockTransactions = []transaction.Transaction{
	{ID: 1, BankAccountFrom: 123, BankAccountTo: 777, Amount: 1.0, CreatedAt: time.Now(), Currency: transaction.RUB},
	{ID: 2, BankAccountFrom: 777, BankAccountTo: 123, Amount: 2.0, CreatedAt: time.Now().AddDate(0, 0, -1), Currency: transaction.RUB},
	{ID: 3, BankAccountFrom: 123, BankAccountTo: 666, Amount: 3.0, CreatedAt: time.Now().AddDate(0, 0, -2), Currency: transaction.RUB},
}

func upFixtures() {
	transaction.AllTransactions = mockTransactions
}

func TestDummyService_List(t *testing.T) {
	upFixtures()

	s := NewDummyTransactionService()
	tx, err := s.List(0, 3)
	require.Empty(t, err, "excepted no err")
	require.Equal(t, 3, len(tx), "wrong number of txs")

	_, err = s.List(10000, 1)
	require.ErrorIs(t, err, ErrOutOfBoundError)

	tx, _ = s.List(1, 1)
	require.Equal(t, mockTransactions[1], tx[0], "got wrong tx")
}

func TestDummyService_Remove(t *testing.T) {
	upFixtures()

	s := NewDummyTransactionService()
	result, _ := s.Remove(1)
	require.Equal(t, true, result, "excepted success")

	tx, _ := s.List(0, 3)
	require.Equal(t, 2, len(tx), "wrong number of txs")

	result, _ = s.Remove(111)
	require.Equal(t, false, result, "excepted error")
}

func TestDummyService_Describe(t *testing.T) {
	upFixtures()

	s := NewDummyTransactionService()
	tx, _ := s.Describe(3)
	require.Equal(t, uint64(3), tx.ID)

	_, err := s.Describe(333)
	require.ErrorIs(t, err, ErrNotExists)
}
