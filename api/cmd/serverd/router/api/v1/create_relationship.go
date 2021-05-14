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

type relationshipCreatorService interface {
	MakeFriend(firstEmail, secondEmail string) (bool, error)
}

// CreateFriend is endpoint to create friend connection between 2 emails addresses.
func (rsv CreateResolver) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var input createFriendInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Description: "Invalid request body"})
		return
	}
	firstEmail := input.Friends[0]
	secondEmail := input.Friends[1]

	result, err := rsv.RelationshipService.MakeFriend(firstEmail, secondEmail)
	if err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "make_friend", Description: err.Error()})
		return
	}

	response.ResponseJson(w, response.Result{Success: result})
}
