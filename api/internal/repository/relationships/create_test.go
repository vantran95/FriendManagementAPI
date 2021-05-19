package relationships

import (
	"errors"
	"regexp"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryImpl_CreateRelationship(t *testing.T) {
	tcs := []struct {
		scenario  string
		input     models.Relationship
		mockErr   error
		expResult bool
		expErr    error
	}{
		{
			scenario:  "success",
			input:     models.Relationship{RequestID: 1, TargetID: 2, Status: "FRIEND"},
			expResult: true,
		},
		{
			scenario:  "invalid data input",
			input:     models.Relationship{RequestID: 1, Status: "FRIEND"},
			expResult: false,
			mockErr:   errors.New("invalid target_id type"),
			expErr:    errors.New("invalid target_id type"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			dbTest, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer dbTest.Close()
			query := regexp.QuoteMeta(`insert into relationships (request_id, target_id, status) values ($1, $2, $3);`)
			if tc.expResult == true {
				mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(tc.input.RequestID, tc.input.TargetID, tc.input.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(tc.input.RequestID, "b@gmail.com", tc.input.Status).
					WillReturnError(errors.New("invalid target_id type"))
			}

			dbMock := &RepositoryImpl{dbTest}
			result, err := dbMock.CreateRelationship(tc.input)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
				assert.NoError(t, tc.expErr, err)
			} else {
				assert.Error(t, tc.expErr, err)
			}
		})
	}
}
