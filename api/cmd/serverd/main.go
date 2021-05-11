package main

import (
	"FriendApi/cmd/serverd/config"
	"FriendApi/cmd/serverd/router"
	"log"
)

func init() {

	//err := godotenv.Load(".env")
	//
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
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
