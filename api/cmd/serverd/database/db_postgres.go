package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

// PostgresDB connect to postgres database
func PostgresDB() (db *sql.DB, err error) {
	dbHost, err := GetEnv("POSTGRES_HOST")
	if err != nil {
		return nil, err
	}
	port, err := GetEnv("POSTGRES_PORT")
	if err != nil {
		return nil, err
	}
	dbPort, _ := strconv.Atoi(port)
	dbUser, err := GetEnv("POSTGRES_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := GetEnv("POSTGRES_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbName, err := GetEnv("POSTGRES_DB")
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	return conn, nil
}
