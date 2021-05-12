package users

import "InternalUserManagement/models"

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
