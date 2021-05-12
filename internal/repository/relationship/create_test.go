package relationship

import (
	"regexp"
	"testing"

	"InternalUserManagement/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateRelationship(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	relationship := models.Relationship{FirstEmailID: 1, SecondEmailID: 2, Status: 0}

	qr := `insert into relationship (first_email_id, second_email_id, status) values ($1, $2, $3);`
	query := regexp.QuoteMeta(qr)
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs(relationship.FirstEmailID, relationship.SecondEmailID, relationship.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// passes the mock to our code
	myDB := NewRelationshipRepository(db)
	result, _ := myDB.CreateRelationship(relationship)

	expected := true
	assert.Equal(t, expected, result)

}
