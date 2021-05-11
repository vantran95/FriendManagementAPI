package v1

import (
	"FriendApi/cmd/serverd/router/restapi/response"
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/exception"
	"encoding/json"
	"net/http"
)

// UserService interface represents the criteria used to retrieve a user service.
type UserService interface {
	GetAllUsers() ([]string, error)
	CreateUser(emailDto dto.EmailDto) (bool, *exception.Exception)
	ExistsByEmail(email string) (bool, error)
	FindUserIdByEmail(email string) (int64, error)
	FindEmailByIds(ids []int64) ([]string, error)
}

// UserAPI stores info to retrieve user service
type UserAPI struct {
	UserService UserService
}

// GetAllUsers retrieve a API to get all users.
func (u UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: users, Count: len(users)}
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

	result, error := u.UserService.CreateUser(emailDto)
	if error != nil {
		response.ErrorResponse(w, http.StatusBadRequest, error.Message)
		return
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}
