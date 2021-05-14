package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	dbUser, dbPassword, dbName := "postgres", "admin", "friend_db"
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)

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
