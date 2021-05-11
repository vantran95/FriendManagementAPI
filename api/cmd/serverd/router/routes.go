package router

import (
	v1 "FriendApi/cmd/serverd/router/restapi/v1"
	users2 "InternalUserManagement/repository/users"
	"InternalUserManagement/service/users"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func initUserController(db *sql.DB) v1.UserAPI {
	userRepository := users2.UserRepositoryImpl{DB: db}
	userService := users.ServiceImpl{Repository: userRepository}
	return v1.UserAPI{UserService: userService}
}

func HandleRequest(db *sql.DB) {
	myRoute := chi.NewRouter()
	userHandle := initUserController(db)
	myRoute.Get("/users", userHandle.GetAllUsers)
	myRoute.Post("/users/create-user", userHandle.CreateUser)

	log.Fatal(http.ListenAndServe(":8082", myRoute))
}
