package main

import (
	"FriendApi/cmd/serverd/config"
	"FriendApi/cmd/serverd/router"
	"github.com/subosito/gotenv"
	"log"
)

func init() {
	gotenv.Load()
}
func main() {
	//dbUser, dbPassword, dbName :=
	//	os.Getenv("POSTGRES_USER"),
	//	os.Getenv("POSTGRES_PASSWORD"),
	//	os.Getenv("POSTGRES_DB")

	dbUser, dbPassword, dbName := "postgres", "admin", "friend_db"
	database, err := config.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()
	router.HandleRequest(database.Conn)
}
