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
	CreateCustomer(name string, email string, status string) (int, string, string, string, error)
	DeleteCustomerById(id string) error
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

func (p *postgres) CreateCustomer(name string, email string, status string) (int, string, string, string, error) {
	insertSQL := `INSERT INTO customers (name, email, status) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, email, status;`
	row := p.conn.QueryRow(insertSQL, name, email, status)

	var id int
	err := row.Scan(&id, &name, &email, &status)
	if err != nil {
		return id, name, email, status, fmt.Errorf("can't insert statement: %w", err)
	}

	return id, name, email, status, nil
}

func (p *postgres) DeleteCustomerById(id string) error {
	stmt, err := p.conn.Prepare("DELETE FROM customers WHERE id=$1;")
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
