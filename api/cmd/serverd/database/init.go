package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	error := godotenv.Load(".env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
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

func initDB(dbType string) (*sql.DB, error) {
	switch dbType {
	case "POSTGRES":
		return PostgresDB()
	default:
		return nil, errors.New("cannot init db")
	}
}
