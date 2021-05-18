package database

import (
	"database/sql"
	"errors"
	"os"
)

const (
	DBTypePostgres = "POSTGRES"
)

// Database stores info for database
type Database struct {
	Conn *sql.DB
}

// Initialize attempts to init a database
func Initialize() (Database, error) {
	dbType := os.Getenv("DB_TYPE")

	conn, err := initDB(dbType)
	if err != nil {
		return Database{}, err
	}

	db := Database{Conn: conn}
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

// initDB attempts to init a database follow data type
func initDB(dbType string) (*sql.DB, error) {
	switch dbType {
	case DBTypePostgres:
		return PostgresDB()
	default:
		return nil, errors.New("cannot init db")
	}
}
