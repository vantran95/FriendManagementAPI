package users

import "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"

// Repository interface represents the criteria used to retrieve a user repository.
type Repository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(email string) (bool, error)
	GetUser(email string) (models.User, error)
}

// ServiceImpl stores info to retrieve user service.
type ServiceImpl struct {
	Repository Repository
}
