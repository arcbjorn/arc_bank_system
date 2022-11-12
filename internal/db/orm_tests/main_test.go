package orm_tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	orm "github.com/arcbjorn/arc_bank_system/internal/db/orm"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5434/arc_bank?sslmode=disable"
)

var testQueries *orm.Queries

func TestMain(m *testing.M) {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = orm.New(connection)

	os.Exit(m.Run())
}
