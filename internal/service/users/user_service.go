package users

import (
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/exception"
	"InternalUserManagement/pkg/utils"
	"errors"
	"net/http"
)

func (s ServiceImpl) GetAllUsers() ([]string, error) {
	return s.Repository.GetAllUsers()
}

// CreateUser .....

// 1. function
// 2. struct
func (s ServiceImpl) CreateUser(emailDto dto.EmailDto) (bool, *exception.Exception) {
	if !utils.IsFormatEmail(emailDto.Email) {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email invalid format"}
	}
	existsByEmail, _ := s.Repository.ExistsByEmail(emailDto.Email)
	if existsByEmail {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email already exists"}
	}
	createUser, _ := s.Repository.CreateUser(emailDto.Email)

	if createUser != true {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Cannot create user"}
	}

	return true, nil
}

func (s ServiceImpl) ExistsByEmail(email string) (bool, error) {
	return s.Repository.ExistsByEmail(email)
}

func (s ServiceImpl) FindUserIdByEmail(email string) (int64, error) {
	id, _ := s.Repository.FindUserIdByEmail(email)
	if id == -1 {
		return -1, errors.New("user not found")
	}
	return id, nil
}
