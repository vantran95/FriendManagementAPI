package relationship

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"InternalUserManagement/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetRelationships(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fromID := int64(1)
	toID := int64(2)

	rows := sqlmock.NewRows([]string{"id", "first_email_id", "second_email_id", "status"}).
		AddRow(1, 1, 2, 0)

	qr := `select x.*
			from relationships x
			where x.first_email_id in (%s, %s)
			and x.second_email_id in (%s, %s);
			`
	query := fmt.Sprintf(
		qr,
		strconv.FormatInt(fromID, 10),
		strconv.FormatInt(toID, 10),
		strconv.FormatInt(fromID, 10),
		strconv.FormatInt(toID, 10))

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// pass the mock to main source
	myDB := NewRelationshipRepository(db)
	results, _ := myDB.GetRelationships(fromID, toID)
	rel := models.Relationship{ID: 1, FirstEmailID: 1, SecondEmailID: 2, Status: 0}
	expected := []models.Relationship{rel}
	assert.Equal(t, expected, results)
}

func TestGetFriendsList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emailID := int64(1)

	rows := sqlmock.NewRows([]string{"id", "email"}).
		AddRow(1, "a@gmail.com").
		AddRow(2, "b@gmail.com")

	qr := `select u.id, u.email
			from users u
         		join relationships r on r.second_email_id = u.id
			where r.first_email_id = %s and r.status = 0
			union
			select u.id, u.email
			from users u
         		join relationships r on r.first_email_id = u.id
			where r.second_email_id = %s and r.status = 0;
			`

	query := fmt.Sprintf(
		qr,
		strconv.FormatInt(emailID, 10),
		strconv.FormatInt(emailID, 10))

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// pass the mock to main source
	myDB := NewRelationshipRepository(db)

	results, _ := myDB.GetFriendsList(emailID)

	user1 := models.User{ID: 1, Email: "a@gmail.com"}
	user2 := models.User{ID: 2, Email: "b@gmail.com"}
	expected := []models.User{user1, user2}
	assert.Equal(t, expected, results)
}
