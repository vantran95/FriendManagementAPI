package users

import (
	"InternalUserManagement/pkg/exception"
	"errors"
	"net/http"
)

// GetAllUsers attempts to get all users
func (s ServiceImpl) GetAllUsers() ([]string, error) {
	return s.Repository.GetAllUsers()
}

// CreateUser attempts to create a user
func (s ServiceImpl) CreateUser(email string) (bool, *exception.Exception) {
	existsByEmail, _ := s.Repository.ExistsByEmail(email)
	if existsByEmail {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email already exists"}
	}
	createUser, _ := s.Repository.CreateUser(email)

	if createUser != true {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Cannot create user"}
	}

	return true, nil
}

// ExistsByEmail attempts to check email is exists
func (s ServiceImpl) ExistsByEmail(email string) (bool, error) {
	return s.Repository.ExistsByEmail(email)
}

// FindUserIdByEmail attempts to find user id by email
func (s ServiceImpl) FindUserIdByEmail(email string) (int64, error) {
	id, _ := s.Repository.FindUserIdByEmail(email)
	if id == -1 {
		return -1, errors.New("user not found")
	}
	return id, nil
}

// FindEmailByIds attempts to find email by user ids.
func (s ServiceImpl) FindEmailByIds(ids []int64) ([]string, error) {
	return s.Repository.FindEmailByIds(ids)
}
