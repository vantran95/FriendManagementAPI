package users

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"

	query := regexp.QuoteMeta(`insert into users (email) values ($1)`)
	mock.ExpectPrepare(query).ExpectExec().WithArgs(email).WillReturnResult(sqlmock.NewResult(1, 1))

	// passes the mock to our code
	myDB := NewUserRepository(db)
	result, _ := myDB.CreateUser(email)

	expected := true
	assert.Equal(t, expected, result)
}
