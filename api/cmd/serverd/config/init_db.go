package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

// Database stores info to retrieve database
type Database struct {
	Conn *sql.DB
}

// Initialize attempts to initialize a database
func Initialize(username, password, dbname string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, dbname)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")

	return db, nil
}
