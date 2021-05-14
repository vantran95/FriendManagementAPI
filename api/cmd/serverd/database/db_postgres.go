package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func PostgresDB() (db *sql.DB, err error) {
	// should refactor to retrieve these from env var
	dbUser, dbPassword, dbName := "postgres", "admin", "friend_db"
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	err = conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")
	return conn, nil
}
