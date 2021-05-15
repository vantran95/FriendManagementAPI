package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	tcs := []struct {
		scenario          string
		mockAPIInput      userCreateInput
		mockServiceInput  string
		mockServiceOutput bool
		mockServiceErr    error
		expResult         interface{}
		expCode           int
	}{
		{
			scenario: "success",
			mockAPIInput: userCreateInput{
				Email: "a@gmail.com",
			},
			mockServiceInput:  "a@gmail.com",
			mockServiceOutput: true,
			expResult: response.Result{
				Success: true,
			},
			expCode: http.StatusOK,
		},
		{
			scenario: "invalid email address",
			mockAPIInput: userCreateInput{
				Email: "invalid", // invalid email
			},

			expResult: response.Error{Status: http.StatusBadRequest, Code: "invalid_request_email", Description: "Invalid email format"},
			expCode:   http.StatusBadRequest,
		},
		{
			scenario: "user already exists",
			mockAPIInput: userCreateInput{
				Email: "a@gmail.com",
			},
			mockServiceInput:  "a@gmail.com",
			mockServiceOutput: false,
			mockServiceErr:    errors.New("email already exists"),
			expResult:         response.Error{Status: http.StatusBadRequest, Code: "create_user", Description: "email already exists"},
			expCode:           http.StatusBadRequest,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			strInput, _ := json.Marshal(tc.mockAPIInput)
			req, err := http.NewRequest(http.MethodPost, "/v1/users/create-user", bytes.NewBuffer(strInput))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			mockResolver := CreateResolver{
				UserService: mockUserCreatorService{
					TestF: t,
					CreateUserInput: struct {
						Input  string
						Output bool
						Err    error
					}{Input: tc.mockServiceInput, Output: tc.mockServiceOutput, Err: tc.mockServiceErr},
				},
			}

			handler := http.HandlerFunc(mockResolver.CreateUser)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)
			assert.Equal(t, string(byteResult), rr.Body.String())
			assert.Equal(t, tc.expCode, rr.Code)
		})
	}
}
