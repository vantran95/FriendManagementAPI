package relationships

import (
	"database/sql"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// NewRelationshipRepository use for create database to write unit test.
func NewRelationshipRepository(db *sql.DB) RepositoryImpl {
	return RepositoryImpl{DB: db}
}

// CreateRelationship attempts to create relationship between 2 email addresses.
func (r RepositoryImpl) CreateRelationship(relationship models.Relationship) (bool, error) {

	query, err := r.DB.Prepare("insert into relationships (first_email_id, second_email_id, status) values ($1, $2, $3);")

	if err != nil {
		return false, err
	}

	_, err = query.Exec(relationship.FirstEmailID, relationship.SecondEmailID, relationship.Status)
	if err != nil {
		return false, err
	}

	return true, nil
}
