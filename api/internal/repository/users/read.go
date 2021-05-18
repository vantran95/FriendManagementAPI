package users

import (
	"database/sql"
	"errors"
	"log"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// GetAllUsers get all users from table users.
func (r RepositoryImpl) GetAllUsers() (*[]models.User, error) {
	result, err := r.DB.Query("select id, email from users")
	if err != nil {
		return nil, err
	}

	var users []models.User

	for result.Next() {
		var id int64
		var email string
		if err := result.Scan(&id, &email); err != nil {
			log.Fatal(err)
		}
		user := models.User{ID: id, Email: email}
		users = append(users, user)
	}
	// Check for no results
	if len(users) == 0 {
		return nil, errors.New("do not have users")
	}
	return &users, nil
}

// GetUser attempts to retrieve a user info
func (r RepositoryImpl) GetUser(email string) (*models.User, error) {
	var id int64
	var userEmail string
	err := r.DB.QueryRow("select id, email from users where email=$1", email).Scan(&id, &userEmail)

	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("user does not exists")
	case err != nil:
		return nil, err
	default:
		user := models.User{ID: id, Email: userEmail}
		return &user, nil
	}
}
