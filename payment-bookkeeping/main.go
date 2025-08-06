package main

import (
	"github.com/fatih/color"
	"github.com/nuxzero/payment-bookkeeping/bookkeeping"
	"github.com/rodaine/table"
)

func main() {
	chargeEntry := bookkeeping.NewChargeEntry(100_000, "USD")

	// Set table styles.
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Set table headers.
	tb := table.New("ID", "AmountSubunits", "AmountCurrency", "Type", "BalanceType", "Kind", "CreatedAt")
	tb.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Set table rows.
	for _, t := range chargeEntry.Transactions {
		tb.AddRow(t.Id, t.AmountSubunits, t.AmountCurrency, t.Type, t.BalanceType, t.Kind, t.CreatedAt)
	}

	tb.Print()
}
