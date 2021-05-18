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

func TestGetFriendsList(t *testing.T) {
	tcs := []struct {
		scenario          string
		mockAPIInput      friendsInput
		mockServiceInput  string
		mockServiceOutput []string
		mockServiceErr    error
		expResult         interface{}
		expCode           int
	}{
		{
			scenario: "success",
			mockAPIInput: friendsInput{
				Email: "a@gmail.com",
			},
			mockServiceInput:  "a@gmail.com",
			mockServiceOutput: []string{"b@gmail.com"},
			expResult:         friendsResponse{Success: true, Friends: []string{"b@gmail.com"}, Count: 1},
			expCode:           http.StatusOK,
		},
		{
			scenario: "user does not exist",
			mockAPIInput: friendsInput{
				Email: "a@gmail.com",
			},
			mockServiceInput: "a@gmail.com",
			mockServiceErr:   errors.New("user does not exists"),
			expResult: response.Error{
				Status:      400,
				Code:        "get_friend_list",
				Description: "user does not exists",
			},
			expCode: http.StatusBadRequest,
		},
		{
			scenario: "user does not have friends",
			mockAPIInput: friendsInput{
				Email: "a@gmail.com",
			},
			mockServiceInput:  "a@gmail.com",
			mockServiceOutput: []string{},
			mockServiceErr:    errors.New("user does not have friend"),
			expResult: response.Error{
				Status:      400,
				Code:        "get_friend_list",
				Description: "user does not have friend",
			},
			expCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			strInput, _ := json.Marshal(tc.mockAPIInput)
			req, err := http.NewRequest(http.MethodPost, "/v1/friend/get-friends-list", bytes.NewBuffer(strInput))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			mockResolver := RetrieveResolver{
				RelationshipService: mockRelationshipRetrieveSrv{
					TestF: t,
					GetFriendsListInput: struct {
						Input  string
						Output []string
						Err    error
					}{Input: tc.mockServiceInput, Output: tc.mockServiceOutput, Err: tc.mockServiceErr},
				},
			}

			handler := http.HandlerFunc(mockResolver.GetFriendsList)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)
			assert.Equal(t, string(byteResult), rr.Body.String())
			assert.Equal(t, tc.expCode, rr.Code)
		})
	}
}

func TestGetCommonFriends(t *testing.T) {
	tcs := []struct {
		scenario                string
		mockAPIInput            commonFriendsInput
		mockServiceRequestInput string
		mockServiceTargetInput  string
		mockServiceOutput       []string
		mockServiceErr          error
		expResult               interface{}
		expCode                 int
	}{
		{
			scenario: "success",
			mockAPIInput: commonFriendsInput{
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
			mockServiceRequestInput: "a@gmail.com",
			mockServiceTargetInput:  "b@gmail.com",
			mockServiceOutput:       []string{"c@gmail.com"},
			expResult:               friendsResponse{Success: true, Friends: []string{"c@gmail.com"}, Count: 1},
			expCode:                 http.StatusOK,
		},
		{
			scenario: "user does not exist",
			mockAPIInput: commonFriendsInput{
				Friends: []string{"aaaa@gmail.com", "b@gmail.com"},
			},
			mockServiceRequestInput: "aaaa@gmail.com",
			mockServiceTargetInput:  "b@gmail.com",
			mockServiceErr:          errors.New("user does not exists"),
			expResult: response.Error{
				Status:      400,
				Code:        "get_common_friends",
				Description: "user does not exists",
			},
			expCode: http.StatusBadRequest,
		},
		{
			scenario: "user does not have common friends",
			mockAPIInput: commonFriendsInput{
				Friends: []string{"a@gmail.com", "b@gmail.com"},
			},
			mockServiceRequestInput: "a@gmail.com",
			mockServiceTargetInput:  "b@gmail.com",
			mockServiceOutput:       []string{},
			mockServiceErr:          errors.New("do not have common friends between two emails"),
			expResult: response.Error{
				Status:      400,
				Code:        "get_common_friends",
				Description: "do not have common friends between two emails",
			},
			expCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			strInput, _ := json.Marshal(tc.mockAPIInput)
			req, err := http.NewRequest(http.MethodPost, "/v1/friend/get-common-friends-list", bytes.NewBuffer(strInput))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			mockResolver := RetrieveResolver{
				RelationshipService: mockRelationshipRetrieveSrv{
					TestF: t,
					GetCommonFriendsInput: struct {
						RequestInput string
						TargetInput  string
						Output       []string
						Err          error
					}{RequestInput: tc.mockServiceRequestInput, TargetInput: tc.mockServiceTargetInput, Output: tc.mockServiceOutput, Err: tc.mockServiceErr},
				},
			}

			handler := http.HandlerFunc(mockResolver.GetCommonFriends)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)
			assert.Equal(t, string(byteResult), rr.Body.String())
			assert.Equal(t, tc.expCode, rr.Code)
		})
	}
}
