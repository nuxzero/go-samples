package bookkeeping

import (
	"reflect"
	"testing"
)

func TestNewChargeEntry(t *testing.T) {
	tests := []struct {
		name     string
		amount   int64
		currency string
		want     Entry
	}{
		{
			name:     "charge 100_000 on card card_123",
			amount:   100_000,
			currency: "USD",
			want: Entry{
				Id:          1,
				Type:        "charge",
				Description: "chrg_123:charge 100000 on card card_123",
				Transactions: []Transaction{
					{
						Id:             1,
						Type:           "Debit",
						BalanceType:    "Charge",
						AmountSubunits: 100_000,
						AmountCurrency: "USD",
						Kind:           "charge.captured.debit",
					},
					{
						Id:             2,
						Type:           "Credit",
						BalanceType:    "Balance",
						AmountSubunits: 96_350,
						AmountCurrency: "USD",
						Kind:           "charge.captured.merchant",
					},
					{
						Id:             3,
						Type:           "Credit",
						BalanceType:    "Balance",
						AmountSubunits: 235,
						AmountCurrency: "USD",
						Kind:           "charge.captured.fee_vat",
					},
					{
						Id:             4,
						Type:           "Credit",
						BalanceType:    "Balance",
						AmountSubunits: 3_650,
						AmountCurrency: "USD",
						Kind:           "charge.captured.fee",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChargeEntry(tt.amount, tt.currency)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChargeEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
