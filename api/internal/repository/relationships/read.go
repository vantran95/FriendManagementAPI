package relationships

import (
	"fmt"
	"strconv"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// GetRelationships attempts to retrieve a relationship through two email ids
func (r RepositoryImpl) GetRelationships(requestID, targetID int64) (*[]models.Relationship, error) {
	stmt := `select x.id, x.request_id, x.target_id, x.status
			from relationships x
			where x.request_id in (%s, %s)
			and x.target_id in (%s, %s);
			`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(requestID, 10),
		strconv.FormatInt(targetID, 10),
		strconv.FormatInt(requestID, 10),
		strconv.FormatInt(targetID, 10))

	results, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	relationships := make([]models.Relationship, 0)
	for results.Next() {
		var id, requestID, targetID int64
		var status string
		err = results.Scan(&id, &requestID, &targetID, &status)
		if err == nil {
			relationships = append(relationships, models.Relationship{ID: id, RequestID: requestID, TargetID: targetID, Status: status})
		}
	}
	return &relationships, nil
}

// GetFriendsList attempts to retrieve a friends list of a email id
func (r RepositoryImpl) GetFriendsList(emailID int64) (*[]models.User, error) {
	qr := `select u.id, u.email
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
		qr,
		strconv.FormatInt(emailID, 10),
		strconv.FormatInt(emailID, 10))

	results, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for results.Next() {
		var id int64
		var email string
		err = results.Scan(&id, &email)
		if err == nil {
			users = append(users, models.User{ID: id, Email: email})
		}
	}
	return &users, nil
}
