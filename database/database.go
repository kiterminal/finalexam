package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Connector interface {
	DeleteById(table string, id string) error
}

type postgres struct {
	conn *sql.DB
}

func Connect() Connector {
	return &postgres{
		conn: db,
	}
}

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

func DirectConn() *sql.DB {
	return db
}

func (p *postgres) DeleteById(table string, id string) error {
	stmt, err := p.conn.Prepare("DELETE FROM " + table + " WHERE id=$1;")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err = stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete: %w", err)
	}

	return nil
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
