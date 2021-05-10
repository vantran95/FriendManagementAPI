package v1

import (
	"FriendApi/cmd/serverd/router/restapi/response"
	"encoding/json"
	"net/http"
)

type UserService interface {
	GetAllUsers() []string
}

type UserAPI struct {
	UserService UserService
}

func (u UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: users, Count: len(users)}
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
