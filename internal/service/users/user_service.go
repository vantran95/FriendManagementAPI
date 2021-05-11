package users

import (
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/exception"
	"InternalUserManagement/pkg/utils"
	"net/http"
)

func (s ServiceImpl) GetAllUsers() []string {
	return s.Repository.GetAllUsers()
}

func (s ServiceImpl) CreateUser(emailDto dto.EmailDto) (bool, *exception.Exception) {
	if !utils.IsFormatEmail(emailDto.Email) {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email invalid format"}
	}
	if s.Repository.ExistsByEmail(emailDto.Email) {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email already exists"}
	}

	createUser := s.Repository.CreateUser(emailDto.Email)

	if createUser != true {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Cannot create user"}
	}

	return true, nil
}

func (s ServiceImpl) ExistsByEmail(email string) bool {
	return s.Repository.ExistsByEmail(email)
}
