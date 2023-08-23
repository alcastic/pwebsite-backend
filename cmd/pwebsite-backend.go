package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alcastic/pwebsite-backend/api"
	"github.com/alcastic/pwebsite-backend/configs"
	"github.com/alcastic/pwebsite-backend/internal/persistence"
	_ "github.com/lib/pq"
)

var config *configs.Config

func init() {
	var err error
	config, err = configs.Configure()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Print("starting")

	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSsslMode,
	)
	db, err := sql.Open(config.DBDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := persistence.NewStore(db)
	server := api.NewServer(store)

	if sErr := server.Start("0.0.0.0:8080"); sErr != nil {
		log.Fatal(sErr)
	}
}
