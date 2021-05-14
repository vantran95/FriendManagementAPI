package users

import (
	"errors"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

type createRepository interface {
	GetUser(email string) (*models.User, error)
	CreateUser(email string) (bool, error)
}

// CreateUser attempts to create a user
func (s ServiceImpl) CreateUser(email string) (bool, error) {
	getUser, _ := s.CreateRepo.GetUser(email)

	if getUser != nil {
		return false, errors.New("user already exist")
	}

	return s.CreateRepo.CreateUser(email)
}
