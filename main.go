package main

import (
	"database/sql"
	"log"

	"github.com/rassulmagauin/simplebank/api"
	db "github.com/rassulmagauin/simplebank/db/sqlc"
	"github.com/rassulmagauin/simplebank/utils"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("unnable to load env variables", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
