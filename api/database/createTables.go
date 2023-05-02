package database

import "fmt"

// createSingleTable creates a single table given an SQL string and panics if anything goest wrong
func createSingleTable(query string) {
	result, err := Db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

// createTables creates all tables in the database if they don't already exist
func createTables() {
	createSingleTable(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMP
		);
	`)
}
