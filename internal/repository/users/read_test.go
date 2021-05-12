package users

import (
	"InternalUserManagement/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

// TestGetAllUser test function get all users.
func TestGetAllUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"
	rows := sqlmock.NewRows([]string{"email"}).AddRow(email)

	query := `select email from user`
	mock.ExpectQuery(query).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	results, _ := myDB.GetAllUsers()

	expected := []string{email}
	assert.Equal(t, expected, results)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"

	rows := sqlmock.NewRows([]string{"id", "email"}).AddRow(1, "a@gmail.com")

	query := regexp.QuoteMeta(`select * from users where email = $1`)
	mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	// pass the mock to main source
	myDB := NewUserRepository(db)
	result, _ := myDB.GetUser(email)

	expected := models.User{ID: 1, Email: "a@gmail.com"}
	assert.Equal(t, expected, result)
}
