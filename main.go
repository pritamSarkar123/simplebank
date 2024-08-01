package main

import (
	"database/sql"
	"log"

	"github.com/golang-projects/simplebank/api"
	db "github.com/golang-projects/simplebank/db/sqlc"
	"github.com/golang-projects/simplebank/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	conn, err := sql.Open(config.DBDriver,config.DBSource)
	if err != nil {
		log.Fatal("can not connect to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
	
}