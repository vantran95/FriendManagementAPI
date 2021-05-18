package users

import "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"

// retrieveRepository interface represents the retrieve repository
type retrieveRepository interface {
	GetUser(email string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
}

// GetAllUsers attempts to get all users
func (s ServiceImpl) GetAllUsers() (*[]models.User, error) {
	return s.RetrieveRepo.GetAllUsers()
}

// GetUser attempts to retrieve user info
func (s ServiceImpl) GetUser(email string) (*models.User, error) {
	return s.RetrieveRepo.GetUser(email)
}
