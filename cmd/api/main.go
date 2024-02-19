package main

import (
	"log"

	"github.com/ithaquaKr/taskManager/config"
	"github.com/ithaquaKr/taskManager/pkg/db/postgres"
)

func main() {
	log.Println("Starting the API Server")

	cfgFile, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Load Config: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Parse Config: %v", err)
	}
	dbConn, err := postgres.NewPostgresConn(cfg)
	if err != nil {
		log.Fatalf("New Postgres Conn: %v", err)
	} else {
		log.Fatalf("PostgreSQL Database connected, Status: %v", dbConn.Stats())
	}
	defer dbConn.Close()
}
