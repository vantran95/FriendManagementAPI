package v1

import (
	"FriendApi/cmd/serverd/router/api/response"
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/exception"
	"encoding/json"
	"net/http"
)

// RelationshipService interface represents the criteria used to retrieve a friend service.
type RelationshipService interface {
	MakeFriend(friendDto dto.FriendDto) (bool, *exception.Exception)
	GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *exception.Exception)
	GetCommonFriends(friendDto dto.FriendDto) ([]string, *exception.Exception)
}

// RelationshipApi stores info to retrieve project friend api
type RelationshipApi struct {
	RelationshipApi RelationshipService
}

// CreateFriend is endpoint to create friend connection between 2 emails addresses.
func (f RelationshipApi) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto

	err := json.NewDecoder(r.Body).Decode(&friendDto)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err1 := f.RelationshipApi.MakeFriend(friendDto)
	if err1 != nil {
		response.ErrorResponse(w, err1.Code, err1.Message)
		return
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

// GetFriendsListByEmail is endpoint to retrieve friends list connected with email
func (f RelationshipApi) GetFriendsListByEmail(w http.ResponseWriter, r *http.Request) {
	var emailDto dto.EmailDto

	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	results, err := f.RelationshipApi.GetFriendsListByEmail(emailDto)
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

// GetCommonFriends is endpoint to retrieve a common friends list.
func (f RelationshipApi) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto

	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	results, err := f.RelationshipApi.GetCommonFriends(friendDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}

	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}
