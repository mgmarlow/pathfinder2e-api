package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_PASS"))

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
