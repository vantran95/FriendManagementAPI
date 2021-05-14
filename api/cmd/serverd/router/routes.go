package router

import (
	"database/sql"
	v1 "github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/v1"
	relationshipRepo "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/repository/relationships"
	userRepo "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/repository/users"
	relationshipService "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/service/relationships"
	userService "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/service/users"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// HandleRequest handle all request route
func HandleRequest(db *sql.DB) {
	routes := chi.NewRouter()

	// Route for user API
	routes.Get("/v1/users", initUserAPIResolver(db).GetAllUsers)
	routes.Post("/v1/users/create-user", initUserAPIResolver(db).CreateUser)

	// Route for relationship API
	routes.Post("/v1/friend/create-friend", initRelationshipAPIResolver(db).CreateFriend)
	routes.Post("/v1/friend/get-friends-list", initRelationshipAPIResolver(db).GetFriendsList)
	routes.Post("/v1/friend/get-common-friends-list", initRelationshipAPIResolver(db).GetCommonFriends)

	log.Fatal(http.ListenAndServe(":8083", routes))
}

func initUserAPIResolver(db *sql.DB) v1.Resolver {
	return v1.Resolver{
		UserSrv: userService.ServiceImpl{
			Repository: userRepo.RepositoryImpl{DB: db},
		},
	}
}

func initRelationshipAPIResolver(db *sql.DB) v1.Resolver {
	return v1.Resolver{
		RelationshipSrv: relationshipService.ServiceImpl{
			UserService: userService.ServiceImpl{
				Repository: userRepo.RepositoryImpl{DB: db},
			},
			Repository: relationshipRepo.RepositoryImpl{DB: db},
		},
	}
}
