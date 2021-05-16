package relationships

import (
	"errors"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceImpl_GetFriendsList(t *testing.T) {
	tcs := []struct {
		scenario                 string
		input                    string
		mockGetUserOutput        *models.User
		mockGetFriendsListOutput *[]models.User
		mockServiceOutput        []string
		mockErr                  error
		expResult                interface{}
		expErr                   error
	}{
		{
			scenario: "success",
			input:    "a@gmail.com",

			mockGetUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetFriendsListOutput: &[]models.User{
				{
					ID:    2,
					Email: "b@gmail.com",
				},
				{
					ID:    3,
					Email: "c@gmail.com",
				},
			},
			mockServiceOutput: []string{"b@gmail.com", "c@gmail.com"},

			expResult: []string{"b@gmail.com", "c@gmail.com"},
		},
		{
			scenario: "user does not exists",
			input:    "a@gmail.com",

			mockGetUserOutput: nil,
			mockErr:           errors.New("user does not exists"),
			expErr:            errors.New("user does not exists"),
		},
		{
			scenario: "user does not have friend",
			input:    "a@gmail.com",

			mockGetUserOutput: nil,
			mockServiceOutput: []string{},
			expResult:         []string{},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockRepo := mockRetrieveRepository{
				TestF: t,
				GetFriendsListInput: struct {
					Input  int64
					Output *[]models.User
					Err    error
				}{Input: 1, Output: tc.mockGetFriendsListOutput, Err: tc.mockErr},
			}
			mockUserServiceRetriever := mockUserServiceRetriever{
				TestF: t,
				GetUserInput: struct {
					Input  string
					Output *models.User
					Err    error
				}{Input: tc.input, Output: tc.mockGetUserOutput, Err: tc.mockErr},
			}

			service := ServiceImpl{
				RetrieveRepo:         mockRepo,
				UserServiceRetriever: mockUserServiceRetriever,
			}

			rs, _ := service.GetFriendsList(tc.input)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
			}
		})
	}
}

func TestServiceImpl_GetCommonFriends(t *testing.T) {
	tcs := []struct {
		scenario          string
		firstInput        string
		secondInput       string
		mockGetFirstFl    []string
		mockGetSecondFL   []string
		mockServiceOutput []string
		mockErr           error
		expResult         interface{}
		expErr            error
	}{
		{
			scenario:        "success",
			firstInput:      "a@gmail.com",
			secondInput:     "b@gmail.com",
			mockGetFirstFl:  []string{"b@gmail.com", "c@gmail.com"},
			mockGetSecondFL: []string{"b@gmail.com", "d@gmail.com"},

			mockServiceOutput: []string{"b@gmail.com"},

			expResult: []string{"b@gmail.com"},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, tc.mockServiceOutput)
			}
		})
	}
}
