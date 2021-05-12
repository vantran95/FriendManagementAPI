package users

import "InternalUserManagement/models"

// GetAllUsers attempts to get all users
func (s ServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}

func (s ServiceImpl) GetUser(email string) (models.User, error) {
	return s.Repository.GetUser(email)
}
