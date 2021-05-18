package users

import (
	"errors"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// retrieveRepository interface represents the retrieve repository
type retrieveRepository interface {
	GetUser(email string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
}

// GetAllUsers attempts to get all users
func (s ServiceImpl) GetAllUsers() ([]models.User, error) {
	users, err := s.RetrieveRepo.GetAllUsers()
	if err != nil {
		return []models.User{}, err
	}
	if len(*users) == 0 {
		return []models.User{}, errors.New("do not have users")
	}
	return *users, nil
}

// GetUser attempts to retrieve user info
func (s ServiceImpl) GetUser(email string) (*models.User, error) {
	return s.RetrieveRepo.GetUser(email)
}
