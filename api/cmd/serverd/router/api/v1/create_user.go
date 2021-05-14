package v1

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
)

// comment here
type userCreateInput struct {
	Email string `json:"email"`
}

type userCreatorService interface {
	CreateUser(email string) (bool, error)
}

// CreateUser retrieve a API to create user.
func (rsv CreateResolver) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input userCreateInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Description: "Invalid request body"})
		return
	}

	email := input.Email
	// Validate email format
	if !isValidEmail(email) {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "invalid_request_email", Description: "Invalid email format"})
		return
	}

	result, error := rsv.UserService.CreateUser(email)
	if error != nil {
		response.ResponseJson(w, response.Error{Status: http.StatusBadRequest, Code: "create_user", Description: error.Error()})
		return
	}

	response.ResponseJson(w, response.Result{Success: result})
}

func isValidEmail(email string) bool {
	const emailRegex = "[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})"
	re, _ := regexp.Compile(emailRegex)

	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !re.MatchString(email) {
		return false
	}
	return true
}
