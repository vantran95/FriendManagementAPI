package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/database"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router"
)

// init attempts to load file .env
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Conn.Close()
	router.HandleRequest(db.Conn)
}
