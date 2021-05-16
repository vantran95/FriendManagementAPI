package v1

import (
	"encoding/json"
	"net/http"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
)

// comment
type getFriendsInput struct {
	Email string `json:"email"`
}

// comment
type getCommonFriendsInput struct {
	Friends []string `json:"friends"`
}

type friendsResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type relationshipRetrieverService interface {
	GetFriendsList(email string) ([]string, error)
	GetCommonFriends(firstEmail, secondEmail string) ([]string, error)
}

// GetFriendsList is endpoint to retrieve friends list connected with email
func (rsv RetrieveResolver) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	var input getFriendsInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Description: "Invalid request body"})
		return
	}

	results, err := rsv.RelationshipService.GetFriendsList(input.Email)
	if err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "get_friend_list", Description: err.Error()})
		return
	}

	if results == nil {
		response.ResponseJson(w, response.Error{Status: http.StatusNotFound, Code: "get_friend_list", Description: "The user doesn't have friends"})
		return
	}

	response.ResponseJson(w, friendsResponse{Success: true, Friends: results, Count: len(results)})
}

// GetCommonFriends is endpoint to retrieve a common friends list.
func (rsv RetrieveResolver) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var input getCommonFriendsInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Description: "Invalid request body"})
		return
	}
	firstEmail := input.Friends[0]
	secondEmail := input.Friends[1]

	results, err := rsv.RelationshipService.GetCommonFriends(firstEmail, secondEmail)
	if err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "get_common_friends", Description: err.Error()})
		return
	}

	if results == nil {
		response.ResponseJson(w, response.Error{Status: http.StatusNotFound, Code: "get_common_friends", Description: "Do not have common friends between two emails"})
		return
	}

	response.ResponseJson(w, friendsResponse{Success: true, Friends: results, Count: len(results)})
}
