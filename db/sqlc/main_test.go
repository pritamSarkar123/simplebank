package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("can not connect to the database")
	}
	testQueries = New(testDB)
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

