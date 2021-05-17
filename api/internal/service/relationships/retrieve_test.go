package relationships

import (
	"errors"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
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

			mockGetFriendsListOutput: nil,
			mockServiceOutput:        []string{},
			expResult:                []string{},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockRepo := mockRetrieveRepository{
				TestF: t,
				GetFriendsListInput: struct {
					Group []struct {
						Input  int64
						Output *[]models.User
						Err    error
					}
				}{Group: []struct {
					Input  int64
					Output *[]models.User
					Err    error
				}{{Input: 1, Output: tc.mockGetFriendsListOutput, Err: tc.mockErr}}},
			}
			mockUserServiceRetriever := mockUserServiceRetriever{
				TestF: t,
				GetUserInput: struct {
					Group []struct {
						Input  string
						Output *models.User
						Err    error
					}
				}{Group: []struct {
					Input  string
					Output *models.User
					Err    error
				}{{Input: tc.input, Output: tc.mockGetUserOutput, Err: tc.mockErr}}},
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
		scenario    string
		firstInput  string
		secondInput string

		mockGetFirstUserOutput  *models.User
		mockGetSecondUserOutput *models.User

		mockGetFirstFriendsListOutput  *[]models.User
		mockGetSecondFriendsListOutput *[]models.User

		mockServiceOutput []string

		mockErr   error
		expResult interface{}
		expErr    error
	}{
		{
			scenario:                "success",
			firstInput:              "a@gmail.com",
			secondInput:             "b@gmail.com",
			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetFirstFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
				{
					ID:    4,
					Email: "d@gmail.com",
				},
			},

			mockGetSecondFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
			},

			mockServiceOutput: []string{"c@gmail.com"},

			expResult: []string{"c@gmail.com"},
		},
		{
			scenario:    "user does not exists",
			firstInput:  "a@gmail.com",
			secondInput: "b@gmail.com",

			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},
			mockErr:                 errors.New("user does not exists"),
			expErr:                  errors.New("user does not exists"),
		},
		{
			scenario:    "do not have common friend",
			firstInput:  "a@gmail.com",
			secondInput: "b@gmail.com",

			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetFirstFriendsListOutput: &[]models.User{
				{
					ID:    4,
					Email: "d@gmail.com",
				},
			},
			mockGetSecondFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
			},
			mockServiceOutput: []string(nil),
			expResult:         []string(nil),
		},
		{
			scenario: "user does not have friend",

			firstInput:              "a@gmail.com",
			secondInput:             "b@gmail.com",
			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetFirstFriendsListOutput: nil,

			mockGetSecondFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
			},
			mockErr: errors.New("user does not have friend"),
			expErr:  errors.New("user does not have friend"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockRepo := mockRetrieveRepository{
				TestF: t,
				GetFriendsListInput: struct {
					Group []struct {
						Input  int64
						Output *[]models.User
						Err    error
					}
				}{
					Group: []struct {
						Input  int64
						Output *[]models.User
						Err    error
					}{
						{
							Input:  1,
							Output: tc.mockGetFirstFriendsListOutput,
							Err:    tc.mockErr,
						},
						{
							Input:  2,
							Output: tc.mockGetSecondFriendsListOutput,
							Err:    tc.mockErr,
						},
					},
				},
			}
			mockUserServiceRetriever := mockUserServiceRetriever{
				TestF: t,
				GetUserInput: struct {
					Group []struct {
						Input  string
						Output *models.User
						Err    error
					}
				}{
					Group: []struct {
						Input  string
						Output *models.User
						Err    error
					}{
						{
							Input:  tc.firstInput,
							Output: tc.mockGetFirstUserOutput,
							Err:    tc.mockErr,
						},
						{
							Input:  tc.secondInput,
							Output: tc.mockGetSecondUserOutput,
							Err:    tc.mockErr,
						},
					},
				}}

			service := ServiceImpl{
				RetrieveRepo:         mockRepo,
				UserServiceRetriever: mockUserServiceRetriever,
			}

			rs, _ := service.GetCommonFriends(tc.firstInput, tc.secondInput)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
			}
		})
	}
}
