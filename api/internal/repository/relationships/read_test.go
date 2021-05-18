package relationships

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryImpl_GetRelationships(t *testing.T) {
	tcs := []struct {
		scenario  string
		requestID int64
		targetID  int64
		mockErr   error
		expResult *[]models.Relationship
		expErr    error
	}{
		{
			scenario:  "success",
			requestID: int64(1),
			targetID:  int64(2),
			expResult: &[]models.Relationship{
				{
					ID:        1,
					RequestID: 1,
					TargetID:  2,
					Status:    "FRIEND",
				},
			},
		},
		{
			scenario:  "no relationship",
			requestID: int64(1),
			targetID:  int64(2),
			expResult: &[]models.Relationship{},
		},
		{
			scenario: "invalid data input",
			targetID: int64(2),
			mockErr:  errors.New("invalid data input"),
			expErr:   errors.New("invalid data input"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			dbTest, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer dbTest.Close()

			stmt := `select x.id, x.request_id, x.target_id, x.status
			from relationships x
			where x.request_id in (%s, %s)
			and x.target_id in (%s, %s);
			`
			query := fmt.Sprintf(
				stmt,
				strconv.FormatInt(tc.requestID, 10),
				strconv.FormatInt(tc.targetID, 10),
				strconv.FormatInt(tc.requestID, 10),
				strconv.FormatInt(tc.targetID, 10))

			if tc.requestID <= 0 {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("invalid data input"))
			} else {
				rows := sqlmock.NewRows([]string{"id", "request_id", "target_id", "status"})
				if len(*tc.expResult) > 0 {
					rows.AddRow(1, tc.requestID, tc.targetID, "FRIEND")
				}
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
			}

			dbMock := &RepositoryImpl{dbTest}
			result, err := dbMock.GetRelationships(tc.requestID, tc.targetID)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
				assert.NoError(t, err)
			} else {
				assert.Error(t, tc.expErr, err)
			}
		})
	}
}

func TestRepositoryImpl_GetFriendsList(t *testing.T) {
	tcs := []struct {
		scenario  string
		emailID   int64
		mockErr   error
		expResult *[]models.User
		expErr    error
	}{
		{
			scenario: "success",
			emailID:  int64(1),
			expResult: &[]models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
			},
		},
		{
			scenario:  "do not have friends list",
			emailID:   int64(1),
			expResult: &[]models.User{},
		},
		{
			scenario: "invalid data input",
			mockErr:  errors.New("invalid data input"),
			expErr:   errors.New("invalid data input"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			dbTest, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer dbTest.Close()

			stmt := `select u.id, u.email
			from users u
        		join relationships r on r.target_id = u.id
			where r.request_id = %s and r.status = 'FRIEND'
			union
			select u.id, u.email
			from users u
        		join relationships r on r.request_id = u.id
			where r.target_id = %s and r.status = 'FRIEND';
			`
			query := fmt.Sprintf(
				stmt,
				strconv.FormatInt(tc.emailID, 10),
				strconv.FormatInt(tc.emailID, 10))
			if tc.emailID <= 0 {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("invalid data input"))
			} else {
				rows := sqlmock.NewRows([]string{"id", "email"})
				if len(*tc.expResult) > 0 {
					for _, v := range *tc.expResult {
						rows.AddRow(v.ID, v.Email)
					}
				}
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
			}

			dbMock := &RepositoryImpl{dbTest}
			result, err := dbMock.GetFriendsList(tc.emailID)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
				assert.NoError(t, err)
			} else {
				assert.Error(t, tc.expErr, err)
			}
		})
	}
}
