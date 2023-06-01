package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect() {
	godotenv.Load()
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	createTables()
	fmt.Println("database connected")
}
