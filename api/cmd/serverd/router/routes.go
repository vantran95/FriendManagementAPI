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

func HandleRequest(db *sql.DB) {
	myRoute := chi.NewRouter()
	myRoute.Get("/users", v1.UserAPI{
		UserService: users.ServiceImpl{
			users2.UserRepositoryImpl{DB: db},
		},
	}.GetAllUsers)

	log.Fatal(http.ListenAndServe(":8082", myRoute))
}
