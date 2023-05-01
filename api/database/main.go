package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Start() {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		5432,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	createTables()
}
