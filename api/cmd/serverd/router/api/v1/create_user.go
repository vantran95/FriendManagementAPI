package v1

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
)

type (
	// userCreateInput stores info to retrieve a json request
	userCreateInput struct {
		Email string `json:"email"`
	}

	// userCreatorService interface represents the criteria used to retrieve a user service
	userCreatorService interface {
		CreateUser(email string) (bool, error)
	}
)

// CreateUser retrieve a API to create user.
func (rsv CreateResolver) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input userCreateInput

	var resErr = response.Error{Status: http.StatusBadRequest}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		resErr.Code = "invalid_request_body"
		resErr.Description = "Invalid request body"
		response.ResponseJson(w, resErr)
		return
	}

	if !input.validate() {
		resErr.Code = "invalid_request_email"
		resErr.Description = "Invalid email format"
		response.ResponseJson(w, resErr)
		return
	}

	_, err = rsv.UserService.CreateUser(input.Email)
	if err != nil {
		resErr.Code = "create_user"
		resErr.Description = err.Error()
		response.ResponseJson(w, resErr)
		return
	}

	response.ResponseJson(w, response.Result{Success: true})
}

// validate attempts to check email is valid
func (u userCreateInput) validate() bool {
	const emailRegex = "[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})"
	re, _ := regexp.Compile(emailRegex)

	return re.MatchString(u.Email)
}
