package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	orm "github.com/arcbjorn/arc_bank_system/internal/db/orm"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5434/arc_bank_test?sslmode=disable"
)

var testQueries *orm.Queries

func cleanUp(testQueries *orm.Queries) {
	err1 := testQueries.DeleteEntries(context.Background())
	if err1 != nil {
		log.Fatal(err1)
	}

	err2 := testQueries.DeleteTransfers(context.Background())
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := testQueries.DeleteAccounts(context.Background())
	if err3 != nil {
		log.Fatal(err3)
	}
}

func TestMain(m *testing.M) {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = orm.New(connection)

	exitCode := m.Run()

	cleanUp(testQueries)

	os.Exit(exitCode)
}
