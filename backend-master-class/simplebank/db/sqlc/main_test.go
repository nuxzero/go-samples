package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// By default, Go will remove all imports that are not used in the code.
	// Since we are not using the pq package directly in our code, Go will remove it.
	// To prevent this, we can use the blank identifier _ to import the package.
	_ "github.com/lib/pq"
	"github.com/nuxzero/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
