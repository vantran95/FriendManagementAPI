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
		response.ErrorResponse(w, error.Code, error.Message)
		return
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

// GetFriendsListByEmail API to get friends list connected with email
func (f FriendApi) GetFriendsListByEmail(w http.ResponseWriter, r *http.Request) {
	var emailDto dto.EmailDto

	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	results, err := f.FriendService.GetFriendsListByEmail(emailDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}

	if len(results) == 0 {
		response.ErrorResponse(w, http.StatusBadRequest, "Can not get friends list")
		return
	}

	res := response.Response{Success: true, Friends: results, Count: len(results)}

	response.SuccessResponse(w, res)
}

func (f FriendApi) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto

	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	results, err := f.FriendService.GetCommonFriends(friendDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}

	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}
