package users

import (
	"InternalUserManagement/pkg/exception"
	"net/http"
)

// CreateUser attempts to create a user
func (s ServiceImpl) CreateUser(email string) (bool, *exception.Exception) {
	getUser, _ := s.Repository.GetUser(email)
	if getUser.Email != "" {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email already exists"}
	}
	createUser, _ := s.Repository.CreateUser(email)

	if createUser != true {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Cannot create user"}
	}

	return true, nil
}
