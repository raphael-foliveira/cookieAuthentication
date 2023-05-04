package database

import (
	"fmt"
	"os"
)

// createTables creates all tables in the database if they don't already exist
func createTables() {
	query, err := os.ReadFile("./database/sql/tables/tables.sql")
	if err != nil {
		fmt.Println("error reading sql file")
		panic(err)
	}
	fmt.Println(string(query))
	result, err := Db.Exec(string(query))
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
