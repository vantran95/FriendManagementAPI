package users

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryImpl_CreateUser(t *testing.T) {
	tcs := []struct {
		scenario             string
		input                string
		mockCreateUserOutput bool
		mockErr              error
		expResult            bool
		expErr               error
	}{
		{
			scenario:             "success",
			input:                "a@mail.com",
			mockCreateUserOutput: true,
			expResult:            true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			//mock.ExpectBegin()
			query := regexp.QuoteMeta(`insert into users (email) values ($1)`)
			mock.ExpectPrepare(query).ExpectExec().WithArgs(tc.input).WillReturnResult(sqlmock.NewResult(1, 1))

			myDB := &RepositoryImpl{db}
			result, _ := myDB.CreateUser(tc.input)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
			}
		})
	}
}
