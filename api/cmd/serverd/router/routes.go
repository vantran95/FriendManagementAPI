package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/v1"
	relationshipRepo "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/repository/relationships"
	userRepo "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/repository/users"
	relationshipService "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/service/relationships"
	userService "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/service/users"
)

// HandleRequest handle all request route
func HandleRequest(db *sql.DB) {
	routes := chi.NewRouter()

	// Route for user API
	routes.Get("/v1/users", initRetrieveResolver(db).GetAllUsers)
	routes.Post("/v1/users/create-user", initCreateResolver(db).CreateUser)

	// Route for relationship API
	routes.Post("/v1/friend/create-friend", initCreateResolver(db).CreateFriend)
	routes.Post("/v1/friend/get-friends-list", initRetrieveResolver(db).GetFriendsList)
	routes.Post("/v1/friend/get-common-friends-list", initRetrieveResolver(db).GetCommonFriends)

	log.Fatal(http.ListenAndServe(":8082", routes))
}

func initRetrieveResolver(db *sql.DB) v1.RetrieveResolver {
	return v1.RetrieveResolver{
		UserService: userService.ServiceImpl{
			CreateRepo:   userRepo.RepositoryImpl{DB: db},
			RetrieveRepo: userRepo.RepositoryImpl{DB: db},
		},
		RelationshipService: relationshipService.ServiceImpl{
			CreateRepo:   relationshipRepo.RepositoryImpl{DB: db},
			RetrieveRepo: relationshipRepo.RepositoryImpl{DB: db},
			UserServiceRetriever: userService.ServiceImpl{
				CreateRepo:   userRepo.RepositoryImpl{DB: db},
				RetrieveRepo: userRepo.RepositoryImpl{DB: db},
			},
		},
	}
}

func initCreateResolver(db *sql.DB) v1.CreateResolver {
	return v1.CreateResolver{
		UserService: userService.ServiceImpl{
			CreateRepo:   userRepo.RepositoryImpl{DB: db},
			RetrieveRepo: userRepo.RepositoryImpl{DB: db},
		},
		RelationshipService: relationshipService.ServiceImpl{
			CreateRepo:   relationshipRepo.RepositoryImpl{DB: db},
			RetrieveRepo: relationshipRepo.RepositoryImpl{DB: db},
			UserServiceRetriever: userService.ServiceImpl{
				CreateRepo:   userRepo.RepositoryImpl{DB: db},
				RetrieveRepo: userRepo.RepositoryImpl{DB: db},
			},
		},
	}
}
