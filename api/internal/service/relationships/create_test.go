package relationships

import (
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceImpl_MakeFriend(t *testing.T) {
	tcs := []struct {
		scenario                string
		firstInput              string
		secondInput             string
		createRelInput          models.Relationship
		mockGetFirstUserOutput  *models.User
		mockGetSecondUserOutput *models.User
		mockGetRelOutput        *[]models.Relationship
		mockServiceOutput       bool
		mockErr                 error
		expResult               interface{}
		expErr                  error
	}{
		{
			scenario:       "success",
			firstInput:     "a@gmail.com",
			secondInput:    "b@gmail.com",
			createRelInput: models.Relationship{FirstEmailID: 1, SecondEmailID: 2, Status: "FRIEND"},

			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},
			mockGetRelOutput: &[]models.Relationship{
				{
					ID:            1,
					FirstEmailID:  1,
					SecondEmailID: 2,
					Status:        "FRIEND",
				},
			},
			mockServiceOutput: true,

			expResult: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockRepo := mockCreateRepository{
				TestF: t,
				GetRelationshipInput: struct {
					FromInput int64
					ToInput   int64
					Output    *[]models.Relationship
					Err       error
				}{FromInput: tc.mockGetFirstUserOutput.ID, ToInput: tc.mockGetSecondUserOutput.ID, Output: tc.mockGetRelOutput, Err: tc.mockErr},
				CreateRelInput: struct {
					Input  models.Relationship
					Output bool
					Err    error
				}{Input: tc.createRelInput, Output: tc.mockServiceOutput, Err: tc.mockErr},
			}
			mockUserService := mockUserServiceRetriever{
				TestF: t,
				GetUserInput: struct {
					Input  string
					Output *models.User
					Err    error
				}{Input: tc.firstInput, Output: tc.mockGetFirstUserOutput, Err: tc.mockErr},
			}

			mockUserService = mockUserServiceRetriever{
				TestF: t,
				GetUserInput: struct {
					Input  string
					Output *models.User
					Err    error
				}{Input: tc.secondInput, Output: tc.mockGetSecondUserOutput, Err: tc.mockErr},
			}

			service := ServiceImpl{
				CreateRepo:           mockRepo,
				UserServiceRetriever: mockUserService,
			}

			rs, _ := service.MakeFriend(tc.firstInput, tc.secondInput)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
			}
		})
	}
}
