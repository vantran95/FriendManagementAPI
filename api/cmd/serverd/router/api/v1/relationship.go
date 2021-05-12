package v1

import (
	"FriendApi/cmd/serverd/router/api/response"
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/utils"
	"encoding/json"
	"net/http"
)

// RelationshipService interface represents the criteria used to retrieve a relationship service.
type RelationshipService interface {
	MakeFriend(firstEmail, secondEmail string) (bool, error)
	GetFriendsList(email string) ([]string, error)
	GetCommonFriends(firstEmail, secondEmail string) ([]string, error)
}

// RelationshipApi stores info to retrieve project relationship api
type RelationshipApi struct {
	RelationshipApi RelationshipService
}

// CreateFriend is endpoint to create friend connection between 2 emails addresses.
func (f RelationshipApi) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto

	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]

	// Validate email format
	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}
	result, err := f.RelationshipApi.MakeFriend(firstEmail, secondEmail)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

// GetFriendsList is endpoint to retrieve friends list connected with email
func (f RelationshipApi) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	var emailDto dto.EmailDto

	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	email := emailDto.Email
	// Validate email format
	if !utils.IsFormatEmail(email) {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}
	results, err := f.RelationshipApi.GetFriendsList(email)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(results) == 0 {
		response.ErrorResponse(w, http.StatusBadRequest, "The user doesn't have friends")
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
	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]

	// Validate email format
	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	results, err := f.RelationshipApi.GetCommonFriends(firstEmail, secondEmail)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}
