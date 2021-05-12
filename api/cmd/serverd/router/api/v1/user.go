package v1

import (
	"encoding/json"
	"net/http"

	"FriendApi/cmd/serverd/router/api/response"
	"InternalUserManagement/models"
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/utils"
)

// UserService interface represents the criteria used to retrieve a user service.
type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(email string) (bool, error)
}

// UserAPI stores info to retrieve user api
type UserAPI struct {
	UserService UserService
}

// GetAllUsers retrieve a API to get all users.
func (u UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserService.GetAllUsers()
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var userEmails []string
	for _, u := range users {
		userEmails = append(userEmails, u.Email)
	}
	res := response.Response{Success: true, Friends: userEmails, Count: len(userEmails)}
	response.SuccessResponse(w, res)
}

// CreateUser retrieve a API to create user.
func (u UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	emailDto := dto.EmailDto{}

	err := json.NewDecoder(r.Body).Decode(&emailDto)

	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	email := emailDto.Email
	// Validate email format
	if !utils.IsFormatEmail(email) {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	result, error := u.UserService.CreateUser(email)
	if error != nil {
		response.ErrorResponse(w, http.StatusBadRequest, error.Error())
		return
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}
