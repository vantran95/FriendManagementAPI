package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	tcs := []struct {
		scenario          string
		mockServiceOutput []models.User
		mockServiceErr    error
		expResult         interface{}
	}{
		{
			scenario: "success",
			mockServiceOutput: []models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
				{
					ID:    2,
					Email: "b@gmail.com",
				},
			},
			expResult: Response{
				Success: true,
				Count:   2,
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
		},
		{
			scenario:          "do not have user",
			mockServiceOutput: []models.User{},
			expResult: response.ErrorResp{
				Status:       http.StatusBadRequest,
				ErrorMessage: "The service does not have any user",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/v1/users", nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			mockResolver := Resolver{
				UserSrv: mockUserService{
					TestF:                 t,
					mockGetAllUsersResult: tc.mockServiceOutput,
					mockError:             tc.mockServiceErr,
				},
			}

			handler := http.HandlerFunc(mockResolver.GetAllUsers)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)

			assert.Equal(t, string(byteResult), rr.Body.String())
			assert.Equal(t, http.StatusOK, rr.Code)
		})
	}
}

func TestCreateUser(t *testing.T) {
	tcs := []struct {
		scenario          string
		mockAPIInput      userCreateInput
		mockServiceInput  string
		mockServiceOutput bool
		mockServiceErr    error
		expResult         interface{}
		isSuccess         bool
	}{
		{
			scenario: "success",
			mockAPIInput: userCreateInput{
				Email: "a@gmail.com",
			},
			mockServiceInput:  "a@gmail.com",
			mockServiceOutput: true,
			expResult: response.SuccessResp{
				Success: true,
			},
		},
		{
			scenario: "invalid email address",
			mockAPIInput: userCreateInput{
				Email: "invalid", // invalid email
			},

			expResult: response.ErrorResp{
				Status:       http.StatusBadRequest,
				ErrorMessage: "Invalid email format",
			},
		},
		{
			scenario: "user already exists",
			mockAPIInput: userCreateInput{
				Email: "a@gmail.com",
			},
			mockServiceInput:  "a@gmail.com",
			mockServiceOutput: false,
			mockServiceErr:    errors.New("email already exists"),
			expResult: response.ErrorResp{
				Status:       http.StatusBadRequest,
				ErrorMessage: "email already exists",
			},
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
			mockResolver := Resolver{
				UserSrv: mockUserService{
					TestF:                t,
					mockCreateUserInput:  tc.mockServiceInput,
					mockCreateUserResult: tc.mockServiceOutput,
					mockError:            tc.mockServiceErr,
				},
			}

			handler := http.HandlerFunc(mockResolver.CreateUser)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)
			assert.Equal(t, string(byteResult), rr.Body.String())
		})
	}
}
