package v1

import (
	"net/http"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

type (
	// userResponse stores info json of user response
	userResponse struct {
		Success bool     `json:"success"`
		Users   []string `json:"users"`
		Count   int      `json:"count"`
	}

	// userRetrieverService interface represents the retrieve service used to retrieve a user info
	userRetrieverService interface {
		GetAllUsers() (*[]models.User, error)
	}
)

// GetAllUsers is endpoint to retrieve users list
func (rsv RetrieveResolver) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := rsv.UserService.GetAllUsers()
	if err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "get_all_users", Description: err.Error()})
		return
	}
	var userEmails []string
	for _, u := range *users {
		userEmails = append(userEmails, u.Email)
	}

	response.ResponseJson(w, userResponse{Success: true, Users: userEmails, Count: len(userEmails)})
}
