package db

import "time"

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
