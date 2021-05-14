package main

import (
	"log"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/database"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router"
)

func main() {
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Conn.Close()
	router.HandleRequest(db.Conn)
}
