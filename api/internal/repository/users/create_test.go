package users

import (
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryImpl_CreateUser(t *testing.T) {
	tcs := []struct {
		scenario  string
		input     string
		mockErr   error
		expResult bool
		expErr    error
	}{
		{
			scenario:  "success",
			input:     "a@mail.com",
			expResult: true,
		},
		{
			scenario:  "email already exists",
			input:     "a@mail.com",
			expResult: false,
			mockErr:   errors.New("email already exists"),
			expErr:    errors.New("email already exists"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			dbTest, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer dbTest.Close()
			query := regexp.QuoteMeta(`insert into users (email) values ($1)`)
			if tc.expResult == true {
				mock.ExpectPrepare(query).ExpectExec().WithArgs(tc.input).WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare(query).ExpectExec().WithArgs(tc.input).WillReturnError(errors.New("email already exists"))
			}
			dbMock := &RepositoryImpl{dbTest}
			result, err := dbMock.CreateUser(tc.input)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
				assert.NoError(t, err)
			}
		})
	}
}
