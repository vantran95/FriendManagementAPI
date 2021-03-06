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

func TestMakeFriend(t *testing.T) {
	tcs := []struct {
		scenario              string
		mockAPIInput          createFriendInput
		mockRequestEmailInput string
		mockTargetEmailInput  string
		mockServiceOutput     bool
		mockServiceErr        error
		expResult             response.Result
		expErr                response.Error
		expCode               int
	}{
		{
			scenario: "success",
			mockAPIInput: createFriendInput{
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
			mockRequestEmailInput: "a@gmail.com",
			mockTargetEmailInput:  "b@gmail.com",
			mockServiceOutput:     true,
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
			mockRequestEmailInput: "aaaaa@gmail.com",
			mockTargetEmailInput:  "b@gmail.com",
			mockServiceOutput:     false,
			mockServiceErr:        errors.New("user does not exists"),
			expErr: response.Error{
				Status:      http.StatusBadRequest,
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
			mockRequestEmailInput: "a@gmail.com",
			mockTargetEmailInput:  "b@gmail.com",
			mockServiceOutput:     false,
			mockServiceErr:        errors.New("already friended"),
			expErr: response.Error{
				Status:      http.StatusBadRequest,
				Code:        "make_friend",
				Description: "already friended",
			},
			expCode: http.StatusBadRequest,
		},
		{
			scenario: "Invalid request body",
			mockAPIInput: createFriendInput{
				Friends: []string{"a@gmail.com"},
			},
			expErr: response.Error{
				Status:      http.StatusBadRequest,
				Code:        "invalid_request_body",
				Description: "Invalid request body",
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
						RequestInput string
						TargetInput  string
						Output       bool
						Err          error
					}{RequestInput: tc.mockRequestEmailInput, TargetInput: tc.mockTargetEmailInput, Output: tc.mockServiceOutput, Err: tc.mockServiceErr},
				},
			}
			handler := http.HandlerFunc(mockResolver.MakeFriend)
			handler.ServeHTTP(rr, req)
			var byteRs []byte
			if tc.expErr.Code != "" {
				byteRs, _ = json.Marshal(tc.expErr)
			} else {
				byteRs, _ = json.Marshal(tc.expResult)
			}
			assert.Equal(t, string(byteRs), rr.Body.String())
			assert.Equal(t, tc.expCode, rr.Code)
		})
	}
}
