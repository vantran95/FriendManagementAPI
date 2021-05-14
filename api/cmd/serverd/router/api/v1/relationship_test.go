package v1

import (
	"bytes"
	"encoding/json"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/cmd/serverd/router/api/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFriend(t *testing.T) {
	var tcs = []struct {
		scenario             string
		mockAPIInput         createFriendInput
		mockFirstEmailInput  string
		mockSecondEmailInput string
		mockServiceOutput    bool
		mockServiceErr       error
		expResult            interface{}
	}{
		{
			scenario: "success",
			mockAPIInput: createFriendInput{
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
			mockFirstEmailInput:  "a@gmail.com",
			mockSecondEmailInput: "b@gmail.com",
			mockServiceOutput:    true,
			expResult: response.SuccessResp{
				Success: true,
			},
		},
		{
			scenario: "user does not exist",
			mockAPIInput: createFriendInput{
				Friends: []string{"aaaaa@gmail.com", "b@gmail.com"},
			},
			//mockFirstEmailInput: "a@gmail.com",
			//mockSecondEmailInput: "b@gmail.com",
			expResult: response.ErrorResp{
				Status:       http.StatusBadRequest,
				ErrorMessage: "cannot get user",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			strInput, _ := json.Marshal(tc.mockAPIInput)
			req, err := http.NewRequest(http.MethodPost, "/v1/friend/create-friend", bytes.NewBuffer(strInput))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			mockResolver := Resolver{
				RelationshipSrv: mockRelationshipService{
					TestF:                     t,
					mockMakeFriendFirstInput:  tc.mockFirstEmailInput,
					mockMakeFriendSecondInput: tc.mockSecondEmailInput,
					mockMakeFriendResult:      tc.mockServiceOutput,
					mockError:                 tc.mockServiceErr,
				},
			}

			handler := http.HandlerFunc(mockResolver.CreateFriend)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)
			assert.Equal(t, string(byteResult), rr.Body.String())
		})
	}
}
