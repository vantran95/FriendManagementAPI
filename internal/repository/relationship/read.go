package relationship

import (
	"InternalUserManagement/models"
	"fmt"
	"strconv"
)

// GetRelationships attempts to retrieve a relationship through two email ids
func (r RepositoryImpl) GetRelationships(fromID, toID int64) ([]models.Relationship, error) {
	stmt := `select x.*
			from relationships x
			where x.first_email_id in (%s, %s)
			and x.second_email_id in (%s, %s);
			`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(fromID, 10),
		strconv.FormatInt(toID, 10),
		strconv.FormatInt(fromID, 10),
		strconv.FormatInt(toID, 10))

	results, err := r.DB.Query(query)
	if err != nil {
		return []models.Relationship{}, err
	}

	var relationships []models.Relationship
	for results.Next() {
		var id, firstEmailID, secondEmailID, status int64
		results.Scan(&id, &firstEmailID, &secondEmailID, &status)
		relationship := models.Relationship{ID: id, FirstEmailID: firstEmailID, SecondEmailID: secondEmailID, Status: status}
		relationships = append(relationships, relationship)
	}
	return relationships, nil
}

// GetFriendsList attempts to retrieve a friends list of a email id
func (r RepositoryImpl) GetFriendsList(emailID int64) ([]models.User, error) {
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

	results, err := r.DB.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	var users []models.User
	for results.Next() {
		var id int64
		var email string
		results.Scan(&id, &email)
		user := models.User{ID: id, Email: email}
		users = append(users, user)
	}

	return users, nil
}
