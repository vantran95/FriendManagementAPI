package relationships

import (
	"regexp"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryImpl_CreateRelationship(t *testing.T) {
	tcs := []struct {
		scenario            string
		input               models.Relationship
		mockCreateRelOutput bool
		mockErr             error
		expResult           bool
		expErr              error
	}{
		{
			scenario:            "success",
			input:               models.Relationship{FirstEmailID: 1, SecondEmailID: 2, Status: "FRIEND"},
			mockCreateRelOutput: true,
			expResult:           true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			query := regexp.QuoteMeta(`insert into relationships (first_email_id, second_email_id, status) values ($1, $2, $3);`)
			mock.ExpectPrepare(query).
				ExpectExec().
				WithArgs(tc.input.FirstEmailID, tc.input.SecondEmailID, tc.input.Status).
				WillReturnResult(sqlmock.NewResult(1, 1))

			myDB := &RepositoryImpl{db}
			result, _ := myDB.CreateRelationship(tc.input)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
			}
		})
	}
}

func TestRepositoryImpl_CreateRelationshipErr(t *testing.T) {
	tcs := []struct {
		scenario            string
		input               models.Relationship
		mockCreateRelOutput bool
		mockErr             error
		expResult           bool
		expErr              error
	}{
		{
			scenario:            "error",
			input:               models.Relationship{FirstEmailID: 1, Status: "FRIEND"},
			mockCreateRelOutput: false,
			expResult:           false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			query := regexp.QuoteMeta(`insert into relationship (first_email_id, second_email_id, status) values ($1, $2, $3);`)
			mock.ExpectPrepare(query).
				ExpectExec().
				WithArgs(tc.input.FirstEmailID, tc.input.SecondEmailID, tc.input.Status).
				WillReturnResult(sqlmock.NewResult(0, 0))

			myDB := &RepositoryImpl{db}
			result, err := myDB.CreateRelationship(tc.input)
			assert.Error(t, err)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
			}
		})
	}
}
