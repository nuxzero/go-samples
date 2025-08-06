package bookkeeping

import (
	"fmt"
	"time"
)

func NewRefund(amount int64, currency string) Entry {
	return Entry{
		Id:          2,
		Type:        "refund",
		Description: fmt.Sprintf("rfnd_123:refund %d on charge chrg_123", amount),
		Transactions: []Transaction{
			{
				Id:             1,
				Type:           "Credit",
				BalanceType:    "Charge",
				AmountSubunits: amount,
				AmountCurrency: currency,
				Kind:           "refund.processed.credit",
				CreatedAt:      time.Now(),
			},
			{
				Id:             2,
				Type:           "Debit",
				BalanceType:    "Balance",
				AmountSubunits: amount - fee(amount) - feeVat(amount),
				AmountCurrency: currency,
				Kind:           "refund.processed.debit",
				CreatedAt:      time.Now(),
			},
		},
	}
}
