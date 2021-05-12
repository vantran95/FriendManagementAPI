package router

import (
	v1 "FriendApi/cmd/serverd/router/api/v1"
	"InternalUserManagement/repository/relationship"
	users2 "InternalUserManagement/repository/users"
	relationship2 "InternalUserManagement/service/relationship"
	"InternalUserManagement/service/users"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// initUserController init a controller for user service
func initUserController(db *sql.DB) v1.UserAPI {
	userRepository := users2.RepositoryImpl{DB: db}
	userService := users.ServiceImpl{Repository: userRepository}
	return v1.UserAPI{UserService: userService}
}

// initRelationshipController init a controller for friend service
func initRelationshipController(db *sql.DB) v1.RelationshipApi {
	userRepository := users2.RepositoryImpl{DB: db}
	userService := users.ServiceImpl{Repository: userRepository}

	relationshipRepository := relationship.RepositoryImpl{DB: db}
	relationshipService := relationship2.ServiceImpl{Repository: relationshipRepository, UserService: userService}

	return v1.RelationshipApi{RelationshipApi: relationshipService}
}

// HandleRequest handle all request route
func HandleRequest(db *sql.DB) {
	myRoute := chi.NewRouter()
	userHandle := initUserController(db)
	myRoute.Get("/users", userHandle.GetAllUsers)
	myRoute.Post("/users/create-user", userHandle.CreateUser)

	// Route for relationship API
	friendHandel := initRelationshipController(db)
	myRoute.Post("/friend/create-friend", friendHandel.CreateFriend)
	myRoute.Post("/friend/get-friends-list", friendHandel.GetFriendsList)
	myRoute.Post("/friend/get-common-friends-list", friendHandel.GetCommonFriends)

	log.Fatal(http.ListenAndServe(":8082", myRoute))
}
