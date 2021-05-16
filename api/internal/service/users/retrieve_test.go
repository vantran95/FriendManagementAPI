package users

import (
	"errors"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceImpl_GetAllUsers(t *testing.T) {
	tcs := []struct {
		scenario              string
		mockGetUserOutput     *models.User
		mockGetAllUsersOutput []models.User
		mockErr               error
		expResult             interface{}
		expErr                error
	}{
		{
			scenario:          "success",
			mockGetUserOutput: nil,
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
			expResult: &[]models.User{
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
			mockGetUserOutput:     nil,
			mockGetAllUsersOutput: nil,
			mockErr:               errors.New("do not have user"),
			expResult:             nil,
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

			assert.Equal(t, tc.expErr, err)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
			}
		})
	}
}
