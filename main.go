package main

import (
	"database/sql"
	"log"

	"github.com/ithaquaKr/taskManager/api"
	db "github.com/ithaquaKr/taskManager/db/sqlc"
	"github.com/ithaquaKr/taskManager/utils"
	_ "github.com/lib/pq"
)

// main is the entry point of the application.
func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to Database:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
