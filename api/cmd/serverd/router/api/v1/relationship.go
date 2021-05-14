package v1

import (
	"encoding/json"
	"net/http"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
)

// comment
type createFriendInput struct {
	Friends []string `json:"friends"`
}

// comment
type getFriendsInput struct {
	Email string `json:"email"`
}

// comment
type getCommonFriendsInput struct {
	Friends []string `json:"friends"`
}

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
func (rsv Resolver) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var input createFriendInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "Invalid request body"})
		return
	}
	firstEmail := input.Friends[0]
	secondEmail := input.Friends[1]

	result, err := rsv.RelationshipSrv.MakeFriend(firstEmail, secondEmail)
	if err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: err.Error()})
		return
	}

	response.ResponseJson(w, r, response.SuccessResp{Success: result})
}

// GetFriendsList is endpoint to retrieve friends list connected with email
func (rsv Resolver) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	var input getFriendsInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "Invalid request body"})
		return
	}

	results, err := rsv.RelationshipSrv.GetFriendsList(input.Email)
	if err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: err.Error()})
		return
	}

	if len(results) == 0 {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "The user doesn't have friends"})
		return
	}

	response.ResponseJson(w, r, response.Response{Success: true, Friends: results, Count: len(results)})
}

// GetCommonFriends is endpoint to retrieve a common friends list.
func (rsv Resolver) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var input getCommonFriendsInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "Invalid request body"})
		return
	}
	firstEmail := input.Friends[0]
	secondEmail := input.Friends[1]

	results, err := rsv.RelationshipSrv.GetCommonFriends(firstEmail, secondEmail)
	if err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: err.Error()})
		return
	}

	response.ResponseJson(w, r, response.Response{Success: true, Friends: results, Count: len(results)})
}
