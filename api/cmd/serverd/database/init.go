package database

import (
	"database/sql"
	"errors"
)

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	// should be choose db type from env var, default is POSTGRES
	dbType := "POSTGRES" // default db type

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

func initDB(dbType string) (*sql.DB, error) {
	switch dbType {
	case "POSTGRES":
		return PostgresDB()
	default:
		return nil, errors.New("cannot init db")
	}
}
