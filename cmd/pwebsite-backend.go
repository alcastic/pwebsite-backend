package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/alcastic/pwebsite-backend/internal/generated/sqlc"
	_ "github.com/lib/pq"
)

var dbusername string
var dbuserpass string
var dbhost string
var dbport string
var dbname string
var dbsslmode string

func init() {
	flag.StringVar(&dbusername, "dbusername", "postgres", "database user name")
	flag.StringVar(&dbuserpass, "dbuserpass", "password", "database user password")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database hostname")
	flag.StringVar(&dbport, "dbport", "5432", "database port")
	flag.StringVar(&dbname, "dbname", "postgres", "database name")
	flag.StringVar(&dbsslmode, "dbsslmode", "disable", "database ssl mode (disable, allow, prefer, require, verify-ca, verify-full)")
}

func main() {
	log.Print("starting")
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", dbusername, dbuserpass, dbhost, dbport, dbname, dbsslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(db)

	lmp := sqlc.ListMessagesParams{PageSize: 2, PageNumber: 0}
	messages, err := queries.ListMessages(context.Background(), lmp)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
