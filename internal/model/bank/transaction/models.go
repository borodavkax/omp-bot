package transaction

import (
	"fmt"
	"time"
)

var AllTransactions = []Transaction{
	{ID: 1, BankAccountFrom: 123, BankAccountTo: 777, Amount: 11.1, CreatedAt: time.Now(), Currency: RUB},
	{ID: 2, BankAccountFrom: 777, BankAccountTo: 123, Amount: 100.0, CreatedAt: time.Now().AddDate(0, 0, -1), Currency: USD},
	{ID: 3, BankAccountFrom: 123, BankAccountTo: 666, Amount: 555.55, CreatedAt: time.Now().AddDate(0, 0, -2), Currency: USD},
	{ID: 4, BankAccountFrom: 777, BankAccountTo: 666, Amount: 1234.5, CreatedAt: time.Now().AddDate(0, 0, -3), Currency: USD},
	{ID: 5, BankAccountFrom: 666, BankAccountTo: 123, Amount: 9999.9, CreatedAt: time.Now().AddDate(0, 0, -4), Currency: RUB},
	{ID: 6, BankAccountFrom: 666, BankAccountTo: 777, Amount: 10.0, CreatedAt: time.Now().AddDate(0, 0, -5), Currency: EUR},
}

type CurrencyCode int64

const (
	RUB CurrencyCode = iota
	USD              = 1
	EUR              = 2
)

type Transaction struct {
	ID              uint64
	BankAccountFrom uint64
	BankAccountTo   uint64
	CreatedAt       time.Time
	Amount          float64
	Currency        CurrencyCode
}

func (m Transaction) String() string {
	return fmt.Sprintf("Transaction: \nID: %v \nFrom: %v \nTo: %v \nAmount: %v \nCreatedAt: %v\nCurrency: %v", m.ID, m.BankAccountFrom, m.BankAccountTo, m.Amount, m.CreatedAt.Format("2006-01-02 15:04:05"), m.Currency)
}

func (s CurrencyCode) String() string {
	switch s {
	case RUB:
		return "RUB"
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	}
	return "unknown"
}
