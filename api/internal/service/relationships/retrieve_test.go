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
		expResult                []string
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

			mockGetUserOutput: &models.User{},
			mockErr:           errors.New("user does not exists"),
			expErr:            errors.New("user does not exists"),
		},
		{
			scenario: "user does not have friend",
			input:    "a@gmail.com",

			mockGetFriendsListOutput: &[]models.User{},
			mockErr:                  errors.New("user does not have friend"),
			expErr:                   errors.New("user does not have friend"),
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
				RetrieveRepo:  mockRepo,
				UserRetriever: mockUserServiceRetriever,
			}

			rs, err := service.GetFriendsList(tc.input)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
				assert.NoError(t, tc.expErr, err)
			}
		})
	}
}

func TestServiceImpl_GetCommonFriends(t *testing.T) {
	tcs := []struct {
		scenario     string
		requestInput string
		targetInput  string

		mockGetRequestUserOutput *models.User
		mockGetTargetUserOutput  *models.User

		mockGetRequestFriendsListOutput *[]models.User
		mockGetTargetFriendsListOutput  *[]models.User

		mockServiceOutput []string

		mockErr   error
		expResult []string
		expErr    error
	}{
		{
			scenario:                 "success",
			requestInput:             "a@gmail.com",
			targetInput:              "b@gmail.com",
			mockGetRequestUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetTargetUserOutput:  &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetRequestFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
				{
					ID:    4,
					Email: "d@gmail.com",
				},
			},

			mockGetTargetFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
			},

			mockServiceOutput: []string{"c@gmail.com"},

			expResult: []string{"c@gmail.com"},
		},
		{
			scenario:     "user does not exists",
			requestInput: "a@gmail.com",
			targetInput:  "b@gmail.com",

			mockGetRequestUserOutput: &models.User{},
			mockGetTargetUserOutput:  &models.User{ID: 2, Email: "b@gmail.com"},
			mockErr:                  errors.New("user does not exists"),
			expErr:                   errors.New("user does not exists"),
		},
		{
			scenario:     "do not have common friend",
			requestInput: "a@gmail.com",
			targetInput:  "b@gmail.com",

			mockGetRequestUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetTargetUserOutput:  &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetRequestFriendsListOutput: &[]models.User{
				{
					ID:    4,
					Email: "d@gmail.com",
				},
			},
			mockGetTargetFriendsListOutput: &[]models.User{
				{
					ID:    3,
					Email: "c@gmail.com",
				},
			},

			mockErr: errors.New("do not have common friends between two emails"),
			expErr:  errors.New("do not have common friends between two emails"),
		},
		{
			scenario: "user does not have friend",

			requestInput:             "a@gmail.com",
			targetInput:              "b@gmail.com",
			mockGetRequestUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetTargetUserOutput:  &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetRequestFriendsListOutput: &[]models.User{},

			mockGetTargetFriendsListOutput: &[]models.User{
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
							Output: tc.mockGetRequestFriendsListOutput,
							Err:    tc.mockErr,
						},
						{
							Input:  2,
							Output: tc.mockGetTargetFriendsListOutput,
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
							Input:  tc.requestInput,
							Output: tc.mockGetRequestUserOutput,
							Err:    tc.mockErr,
						},
						{
							Input:  tc.targetInput,
							Output: tc.mockGetTargetUserOutput,
							Err:    tc.mockErr,
						},
					},
				}}

			service := ServiceImpl{
				RetrieveRepo:  mockRepo,
				UserRetriever: mockUserServiceRetriever,
			}

			rs, err := service.GetCommonFriends(tc.requestInput, tc.targetInput)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
				assert.NoError(t, tc.expErr, err)
			}
		})
	}
}
