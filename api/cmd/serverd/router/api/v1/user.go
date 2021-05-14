package v1

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// comment here
type userCreateInput struct {
	Email string `json:"email"`
}

// UserService interface represents the criteria used to retrieve a user service.
type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(email string) (bool, error)
}

func (rsv Resolver) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := rsv.UserSrv.GetAllUsers()
	if err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: err.Error()})
		return
	}
	var userEmails []string
	for _, u := range users {
		userEmails = append(userEmails, u.Email)
	}

	if len(userEmails) == 0 {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "The service does not have any user"})
		return
	}

	response.ResponseJson(w, r, response.Response{Success: true, Friends: userEmails, Count: len(userEmails)})
}

// CreateUser retrieve a API to create user.
func (rsv Resolver) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input userCreateInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "Invalid request body"})
		return
	}

	email := input.Email
	// Validate email format
	if !isValidEmail(email) {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: "Invalid email format"})
		return
	}

	result, error := rsv.UserSrv.CreateUser(email)
	if error != nil {
		response.ResponseJson(w, r, response.ErrorResp{Status: http.StatusBadRequest, ErrorMessage: error.Error()})
		return
	}

	response.ResponseJson(w, r, response.SuccessResp{Success: result})
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
