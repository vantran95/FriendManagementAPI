package v1

import (
	"encoding/json"
	"net/http"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
)

type (
	// createFriendInput stores info to retrieve a json of friend response
	createFriendInput struct {
		Friends []string `json:"friends"`
	}

	// relationshipCreatorService interface represents the create relationship service
	relationshipCreatorService interface {
		MakeFriend(requestEmail, targetEmail string) (bool, error)
	}
)

// MakeFriend is endpoint to make friend connection between 2 emails addresses.
func (rsv CreateResolver) MakeFriend(w http.ResponseWriter, r *http.Request) {
	var input createFriendInput
	var resErr = response.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Description: "Invalid request body"}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.ResponseJson(w, resErr)
		return
	}
	if len(input.Friends) < 2 {
		response.ResponseJson(w, resErr)
		return
	}
	requestEmail := input.Friends[0]
	targetEmail := input.Friends[1]

	result, err := rsv.RelationshipService.MakeFriend(requestEmail, targetEmail)
	if err != nil {
		resErr.Code = "make_friend"
		resErr.Description = err.Error()
		response.ResponseJson(w, resErr)
		return
	}
	response.ResponseJson(w, response.Result{Success: result})
}
