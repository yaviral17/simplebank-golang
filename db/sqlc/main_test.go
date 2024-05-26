package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:toor@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(
	m *testing.M,
) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Printf("cannot connect to db: %v", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}
