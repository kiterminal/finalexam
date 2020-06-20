package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("connect to database error", err)
	}

	if err = createCustomerTable(); err != nil {
		log.Fatal(err)
	}
}

func Conn() *sql.DB {
	return db
}

func createCustomerTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`

	_, err := db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("can't create customers table: %w", err)
	}

	return nil
}
