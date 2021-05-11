package v1

import (
	"FriendApi/cmd/serverd/router/restapi/response"
	"InternalUserManagement/pkg/dto"
	"encoding/json"
	"net/http"
)

type FriendApi struct {
	FriendService FriendService
}

// CreateFriend API to create friend connection between 2 emails addresses.
func (f FriendApi) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto

	err := json.NewDecoder(r.Body).Decode(&friendDto)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, error := f.FriendService.CreateFriend(friendDto)
	if error != nil {
		response.ErrorResponse(w, error.Code, error.Error())
		return
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}
