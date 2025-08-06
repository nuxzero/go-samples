package bookkeeping

import (
	"fmt"
	"time"
)

const (
	FeeRate = 3.65
	FeeVat  = 7
)

type Entry struct {
	Id           int64
	Description  string
	Type         string
	Transactions []Transaction
	CreatedAt    time.Time
}

type Transaction struct {
	Id             int64
	AmountSubunits int64
	AmountCurrency string
	Type           string
	BalanceType    string
	Kind           string
	CreatedAt      time.Time
}

type Balance struct {
	Id            int64
	TotalSubunits int64
	Currency      string
	CreatedAt     time.Time
}

func NewChargeEntry(amount int64, currency string) Entry {
	return Entry{
		Id:          1,
		Type:        "charge",
		Description: fmt.Sprintf("chrg_123:charge %d on card card_123", amount),
		Transactions: []Transaction{
			{
				Id:             1,
				Type:           "Debit",
				BalanceType:    "Charge",
				AmountSubunits: amount,
				AmountCurrency: currency,
				Kind:           "charge.captured.debit",
				CreatedAt:      time.Now(),
			},
			{
				Id:             2,
				Type:           "Credit",
				BalanceType:    "Balance",
				AmountSubunits: amount - fee(amount) - feeVat(amount),
				AmountCurrency: currency,
				Kind:           "charge.captured.merchant",
				CreatedAt:      time.Now(),
			},
			{
				Id:             3,
				Type:           "Credit",
				BalanceType:    "Balance",
				AmountSubunits: feeVat(amount),
				AmountCurrency: currency,
				Kind:           "charge.captured.fee_vat",
				CreatedAt:      time.Now(),
			},
			{
				Id:             4,
				Type:           "Credit",
				BalanceType:    "Balance",
				AmountSubunits: fee(amount),
				AmountCurrency: currency,
				Kind:           "charge.captured.fee",
				CreatedAt:      time.Now(),
			},
		},
		CreatedAt: time.Now(),
	}
}

func fee(amount int64) int64 {
	return int64(float64(amount) * FeeRate / 100)
}

func feeVat(amount int64) int64 {
	return fee(amount) * FeeVat / 100
}
