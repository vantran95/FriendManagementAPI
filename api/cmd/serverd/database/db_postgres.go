package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// PostgresDB connect to postgres database
func PostgresDB() (db *sql.DB, err error) {
	error := godotenv.Load(".env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
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
