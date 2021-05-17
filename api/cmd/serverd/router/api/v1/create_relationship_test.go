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

func TestCreateFriend(t *testing.T) {
	tcs := []struct {
		scenario             string
		mockAPIInput         createFriendInput
		mockFirstEmailInput  string
		mockSecondEmailInput string
		mockServiceOutput    bool
		mockServiceErr       error
		expResult            interface{}
		expCode              int
	}{
		{
			scenario: "success",
			mockAPIInput: createFriendInput{
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
			mockFirstEmailInput:  "a@gmail.com",
			mockSecondEmailInput: "b@gmail.com",
			mockServiceOutput:    true,
			expResult: response.Result{
				Success: true,
			},
			expCode: http.StatusOK,
		},
		{
			scenario: "user does not exist",
			mockAPIInput: createFriendInput{
				Friends: []string{"aaaaa@gmail.com", "b@gmail.com"},
			},
			mockFirstEmailInput:  "aaaaa@gmail.com",
			mockSecondEmailInput: "b@gmail.com",
			mockServiceOutput:    false,
			mockServiceErr:       errors.New("user does not exists"),
			expResult: response.Error{
				Status:      400,
				Code:        "make_friend",
				Description: "user does not exists",
			},
			expCode: http.StatusBadRequest,
		},
		{
			scenario: "already friended",
			mockAPIInput: createFriendInput{
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
			mockFirstEmailInput:  "a@gmail.com",
			mockSecondEmailInput: "b@gmail.com",
			mockServiceOutput:    false,
			mockServiceErr:       errors.New("already friended"),
			expResult: response.Error{
				Status:      400,
				Code:        "make_friend",
				Description: "already friended",
			},
			expCode: http.StatusBadRequest,
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
			mockResolver := CreateResolver{
				RelationshipService: mockRelationshipCreatorSrv{
					TestF: t,
					MakeFriendInput: struct {
						FirstInput  string
						SecondInput string
						Output      bool
						Err         error
					}{FirstInput: tc.mockFirstEmailInput, SecondInput: tc.mockSecondEmailInput, Output: tc.mockServiceOutput, Err: tc.mockServiceErr},
				},
			}

			handler := http.HandlerFunc(mockResolver.CreateFriend)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)
			assert.Equal(t, string(byteResult), rr.Body.String())
			assert.Equal(t, tc.expCode, rr.Code)
		})
	}
}
