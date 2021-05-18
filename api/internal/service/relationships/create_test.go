package relationships

import (
	"errors"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestServiceImpl_MakeFriend(t *testing.T) {
	tcs := []struct {
		scenario       string
		requestInput   string
		targetInput    string
		createRelInput models.Relationship

		mockGetRequestUserOutput *models.User
		mockGetTargetUserOutput  *models.User

		mockGetRelOutput *[]models.Relationship

		mockServiceOutput bool
		mockErr           error
		expResult         bool
		expErr            error
	}{
		{
			scenario:     "success",
			requestInput: "a@gmail.com",
			targetInput:  "b@gmail.com",

			createRelInput: models.Relationship{RequestID: 1, TargetID: 2, Status: "FRIEND"},

			mockGetRequestUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetTargetUserOutput:  &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetRelOutput:  &[]models.Relationship{},
			mockServiceOutput: true,
			expResult:         true,
		},
		{
			scenario: "user do not exists",

			requestInput: "a@gmail.com",
			targetInput:  "b@gmail.com",

			mockGetRequestUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetTargetUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},

			mockErr: errors.New("user does not exists"),
			expErr:  errors.New("user does not exists"),
		},
		{
			scenario:     "already friended",
			requestInput: "a@gmail.com",
			targetInput:  "b@gmail.com",

			createRelInput: models.Relationship{RequestID: 1, TargetID: 2, Status: "FRIEND"},

			mockGetRequestUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetTargetUserOutput:  &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetRelOutput: &[]models.Relationship{
				{
					ID:        1,
					RequestID: 1,
					TargetID:  2,
					Status:    "FRIEND",
				},
			},
			mockServiceOutput: false,

			expResult: false,

			mockErr: errors.New("already friended"),
			expErr:  errors.New("already friended"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			mockUserService := mockUserServiceRetriever{
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

			mockRetrieveRepo := mockRetrieveRepository{
				TestF: t,
				GetRelationshipInput: struct {
					RequestInput int64
					TargetInput  int64
					Output       *[]models.Relationship
					Err          error
				}{RequestInput: tc.mockGetRequestUserOutput.ID, TargetInput: tc.mockGetTargetUserOutput.ID, Output: tc.mockGetRelOutput, Err: tc.mockErr},
			}

			mockCreateRepo := mockCreateRepository{
				TestF: t,
				CreateRelInput: struct {
					Input  models.Relationship
					Output bool
					Err    error
				}{Input: tc.createRelInput, Output: tc.mockServiceOutput, Err: tc.mockErr},
			}

			service := ServiceImpl{
				CreateRepo:    mockCreateRepo,
				RetrieveRepo:  mockRetrieveRepo,
				UserRetriever: mockUserService,
			}

			rs, err := service.MakeFriend(tc.requestInput, tc.targetInput)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
				assert.NoError(t, err)
			}
		})
	}
}
