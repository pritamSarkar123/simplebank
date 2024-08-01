package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/golang-projects/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M){
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	testDB, err = sql.Open(config.DBDriver,config.DBSource)
	if err != nil {
		log.Fatal("can not connect to the database")
	}
	testQueries = New(testDB)
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

