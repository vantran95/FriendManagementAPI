package users

import (
	"errors"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestServiceImpl_GetAllUsers(t *testing.T) {
	tcs := []struct {
		scenario              string
		mockGetAllUsersOutput []models.User
		mockErr               error
		expResult             []models.User
		expErr                error
	}{
		{
			scenario: "success",
			mockGetAllUsersOutput: []models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
				{
					ID:    2,
					Email: "b@gmail.com",
				},
			},
			expResult: []models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
				{
					ID:    2,
					Email: "b@gmail.com",
				},
			},
		},
		{
			scenario:              "do not have user",
			mockGetAllUsersOutput: []models.User{},
			mockErr:               errors.New("do not have user"),
			expResult:             []models.User{},
			expErr:                errors.New("do not have user"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockRepo := mockRetrieveRepository{
				TestF: t,
				GetAllUsersInput: struct {
					Output *[]models.User
					Err    error
				}{Output: &tc.mockGetAllUsersOutput, Err: tc.mockErr},
			}
			service := ServiceImpl{
				RetrieveRepo: mockRepo,
			}
			rs, err := service.GetAllUsers()
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
				assert.NoError(t, tc.expErr, err)
			} else {
				assert.Error(t, tc.expErr, err)
			}
		})
	}
}

func TestServiceImpl_GetUser(t *testing.T) {
	tcs := []struct {
		scenario          string
		input             string
		mockGetUserOutput *models.User
		mockErr           error
		expResult         *models.User
		expErr            error
	}{
		{
			scenario:          "success",
			input:             "a@gmail.com",
			mockGetUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			expResult:         &models.User{ID: 1, Email: "a@gmail.com"},
		},
		{
			scenario:          "do not have user",
			input:             "a@gmail.com",
			mockGetUserOutput: &models.User{},
			mockErr:           errors.New("do not have user"),
			expErr:            errors.New("do not have user"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockRepo := mockRetrieveRepository{
				TestF: t,
				GetUserInput: struct {
					Input  string
					Output *models.User
					Err    error
				}{Input: tc.input, Output: tc.mockGetUserOutput, Err: tc.mockErr},
			}
			service := ServiceImpl{
				RetrieveRepo: mockRepo,
			}
			rs, err := service.GetUser(tc.input)

			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
				assert.NoError(t, tc.expErr, err)
			} else {
				assert.Error(t, tc.expErr, err)
			}
		})
	}
}
