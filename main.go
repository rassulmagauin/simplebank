package main

import (
	"database/sql"
	"log"

	"github.com/rassulmagauin/simplebank/api"
	db "github.com/rassulmagauin/simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
