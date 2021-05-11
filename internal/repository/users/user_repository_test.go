package users

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strconv"
	"strings"
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

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"

	query := regexp.QuoteMeta(`insert into user_management (email) values ($1)`)
	mock.ExpectPrepare(query).ExpectExec().WithArgs(email).WillReturnResult(sqlmock.NewResult(1, 1))

	// passes the mock to our code
	myDB := NewUserRepository(db)
	result, _ := myDB.CreateUser(email)

	expected := true
	assert.Equal(t, expected, result)
}

func TestExistsByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	query := regexp.QuoteMeta(`select id from user_management where email = $1`)
	mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	result, _ := myDB.ExistsByEmail(email)

	expected := true
	assert.Equal(t, expected, result)
}

func TestFindUserIdByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	query := regexp.QuoteMeta(`select id from user_management where email = $1`)
	mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	result, _ := myDB.FindUserIdByEmail(email)

	expected := int64(1)
	assert.Equal(t, expected, result)
}

func TestFindEmailByIds(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ids := []int64{1, 2, 3}
	rows := sqlmock.NewRows([]string{"email"}).
		AddRow("a@gmail.com").
		AddRow("b@gmail.com").
		AddRow("c@gmail.com")

	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.email from user_management x where x.id in (%s);`
	query := fmt.Sprintf(stmt, strings.Join(strIds, ","))
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	results, _ := myDB.FindEmailByIds(ids)

	expected := []string{"a@gmail.com", "b@gmail.com", "c@gmail.com"}
	assert.Equal(t, expected, results)
}
