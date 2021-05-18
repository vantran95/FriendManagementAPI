package v1

import (
	"encoding/json"
	"net/http"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
)

type (
	// getFriendsInput stores info to retrieve friend input json
	friendsInput struct {
		Email string `json:"email"`
	}

	// getCommonFriendsInput stores info to retrieve common friend input json
	commonFriendsInput struct {
		Friends []string `json:"friends"`
	}

	// friendsResponse stores info json of friend response
	friendsResponse struct {
		Success bool     `json:"success"`
		Friends []string `json:"friends"`
		Count   int      `json:"count"`
	}

	// relationshipRetrieverService interface represents the retrieve service used to retrieve relationships
	relationshipRetrieverService interface {
		GetFriendsList(email string) ([]string, error)
		GetCommonFriends(requestEmail, targetEmail string) ([]string, error)
	}
)

// GetFriendsList is endpoint to retrieve friends list connected with email
func (rsv RetrieveResolver) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	var input friendsInput
	var resErr = response.Error{Status: http.StatusBadRequest}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		resErr.Code = "invalid_request_body"
		resErr.Description = "Invalid request body"
		response.ResponseJson(w, resErr)
		return
	}

	results, err := rsv.RelationshipService.GetFriendsList(input.Email)
	if err != nil {
		resErr.Code = "get_friend_list"
		resErr.Description = err.Error()
		response.ResponseJson(w, resErr)
		return
	}

	response.ResponseJson(w, friendsResponse{Success: true, Friends: results, Count: len(results)})
}

// GetCommonFriends is endpoint to retrieve a common friends list.
func (rsv RetrieveResolver) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var input commonFriendsInput
	var resErr = response.Error{Status: http.StatusBadRequest}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		resErr.Code = "invalid_request_body"
		resErr.Description = "Invalid request body"
		response.ResponseJson(w, resErr)
		return
	}
	requestEmail := input.Friends[0]
	targetEmail := input.Friends[1]

	results, err := rsv.RelationshipService.GetCommonFriends(requestEmail, targetEmail)
	if err != nil {
		resErr.Code = "get_common_friends"
		resErr.Description = err.Error()
		response.ResponseJson(w, resErr)
		return
	}

	response.ResponseJson(w, friendsResponse{Success: true, Friends: results, Count: len(results)})
}
