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
}

type UserAPI struct {
	UserService UserService
}

func (u UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: users, Count: len(users)}
	responseWithJSON(w, http.StatusOK, res)
}

func (u UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	emailDto := dto.EmailDto{}

	err := json.NewDecoder(r.Body).Decode(&emailDto)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, error := u.UserService.CreateUser(emailDto)
	if error != nil {
		responseWithError(w, http.StatusBadRequest, error.Error())
	}

	res := response.Success{Success: result}
	responseWithJSON(w, http.StatusOK, res)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"Error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
