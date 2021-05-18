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
		firstInput     string
		secondInput    string
		createRelInput models.Relationship

		mockGetFirstUserOutput  *models.User
		mockGetSecondUserOutput *models.User

		mockGetRelOutput *[]models.Relationship

		mockServiceOutput bool
		mockErr           error
		expResult         interface{}
		expErr            error
	}{
		{
			scenario:    "success",
			firstInput:  "a@gmail.com",
			secondInput: "b@gmail.com",

			createRelInput: models.Relationship{FirstEmailID: 1, SecondEmailID: 2, Status: "FRIEND"},

			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},

			mockGetRelOutput:  nil,
			mockServiceOutput: true,
			expResult:         true,
		},
		{
			scenario: "user do not exists",

			firstInput:  "a@gmail.com",
			secondInput: "b@gmail.com",

			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 1, Email: "a@gmail.com"},

			mockErr: errors.New("user does not exists"),
			expErr:  errors.New("user does not exists"),
		},
		{
			scenario:    "already friended",
			firstInput:  "a@gmail.com",
			secondInput: "b@gmail.com",

			createRelInput: models.Relationship{FirstEmailID: 1, SecondEmailID: 2, Status: "FRIEND"},

			mockGetFirstUserOutput:  &models.User{ID: 1, Email: "a@gmail.com"},
			mockGetSecondUserOutput: &models.User{ID: 2, Email: "b@gmail.com"},
			//mockGetRelOutput: nil,

			mockGetRelOutput: &[]models.Relationship{
				{
					ID:            1,
					FirstEmailID:  1,
					SecondEmailID: 2,
					Status:        "FRIEND",
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

			mockRetrieveRepo := mockRetrieveRepository{
				TestF: t,
				GetRelationshipInput: struct {
					FromInput int64
					ToInput   int64
					Output    *[]models.Relationship
					Err       error
				}{FromInput: tc.mockGetFirstUserOutput.ID, ToInput: tc.mockGetSecondUserOutput.ID, Output: tc.mockGetRelOutput, Err: tc.mockErr},
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

			rs, err := service.MakeFriend(tc.firstInput, tc.secondInput)

			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, rs)
				assert.NoError(t, err)
			}
		})
	}
}
