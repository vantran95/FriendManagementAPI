package users

import (
	"errors"
	"regexp"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryImpl_GetUser(t *testing.T) {
	tcs := []struct {
		scenario  string
		input     string
		mockErr   error
		expResult *models.User
		expErr    error
	}{
		{
			scenario:  "success",
			input:     "a@mail.com",
			expResult: &models.User{ID: 1, Email: "a@gmail.com"},
		},
		{
			scenario: "user does not exists",
			input:    "a@mail.com",

			mockErr: errors.New("user does not exists"),
			expErr:  errors.New("user does not exists"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			dbTest, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer dbTest.Close()
			rows := sqlmock.NewRows([]string{"id", "email"})
			if tc.expResult != nil {
				rows.AddRow(tc.expResult.ID, tc.expResult.Email)
			}
			query := regexp.QuoteMeta(`select id, email from users where email=$1`)
			mock.ExpectQuery(query).WithArgs(tc.input).WillReturnRows(rows)

			dbMock := &RepositoryImpl{dbTest}
			result, err := dbMock.GetUser(tc.input)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
				assert.NoError(t, tc.expErr, err)
			}
		})
	}
}

func TestRepositoryImpl_GetAllUsers(t *testing.T) {
	tcs := []struct {
		scenario  string
		mockErr   error
		expResult *[]models.User
		expErr    error
	}{
		{
			scenario: "success",
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
			scenario:  "do not have users",
			expResult: &[]models.User{},
			mockErr:   errors.New("do not have users"),
			expErr:    errors.New("do not have users"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			dbTest, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer dbTest.Close()
			rows := sqlmock.NewRows([]string{"id", "email"})
			if len(*tc.expResult) > 0 {
				for _, v := range *tc.expResult {
					rows.AddRow(v.ID, v.Email)
				}
			}
			query := regexp.QuoteMeta(`select id, email from users`)
			mock.ExpectQuery(query).WillReturnRows(rows)

			dbMock := &RepositoryImpl{dbTest}
			result, err := dbMock.GetAllUsers()
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
				assert.NoError(t, tc.expErr, err)
			}
		})
	}
}
