package users

import "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"

// GetAllUsers attempts to get all users
func (s ServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}

func (s ServiceImpl) GetUser(email string) (models.User, error) {
	return s.Repository.GetUser(email)
}
