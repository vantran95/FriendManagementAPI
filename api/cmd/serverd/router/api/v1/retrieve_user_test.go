package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	tcs := []struct {
		scenario          string
		mockServiceOutput []models.User
		mockServiceErr    error
		expResult         userResponse
	}{
		{
			scenario: "success",
			mockServiceOutput: []models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
				{
					ID:    2,
					Email: "b@gmail.com",
				},
			},
			expResult: userResponse{
				Success: true,
				Count:   2,
				Users:   []string{"a@gmail.com", "b@gmail.com"},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/v1/users", nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			mockResolver := RetrieveResolver{
				UserService: mockUserRetrieverService{
					TestF: t,
					GetAllUsersInput: struct {
						Output *[]models.User
						Err    error
					}{Output: &tc.mockServiceOutput, Err: tc.mockServiceErr},
				},
			}

			handler := http.HandlerFunc(mockResolver.GetAllUsers)
			handler.ServeHTTP(rr, req)

			byteResult, _ := json.Marshal(tc.expResult)

			assert.Equal(t, string(byteResult), rr.Body.String())
			assert.Equal(t, http.StatusOK, rr.Code)
		})
	}
}
