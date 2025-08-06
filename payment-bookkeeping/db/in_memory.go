package db

// Create in memory database
type Database struct {
	entries  []Entry
	balances []Balance
}

// Create a new database
func New() *Database {
	return &Database{}
}

// Add an entry to the database
func (d *Database) AddEntry(e Entry) {
	d.entries = append(d.entries, e)
}

// Add a balance to the database
func (d *Database) AddBalance(b Balance) {
	d.balances = append(d.balances, b)
}
