package v1

import (
	"FriendApi/cmd/serverd/router/restapi/response"
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/exception"
	"encoding/json"
	"net/http"
)

type UserService interface {
	GetAllUsers() []string
	CreateUser(emailDto dto.EmailDto) (bool, *exception.Exception)
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) (int64, error)
}

type UserAPI struct {
	UserService UserService
}

func (u UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: users, Count: len(users)}
	response.SuccessResponse(w, res)
}

func (u UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	emailDto := dto.EmailDto{}

	err := json.NewDecoder(r.Body).Decode(&emailDto)

	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, error := u.UserService.CreateUser(emailDto)
	if error != nil {
		response.ErrorResponse(w, http.StatusBadRequest, error.Error())
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}
