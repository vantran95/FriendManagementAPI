package users

import (
	"database/sql"
	"errors"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"log"
)

// GetAllUsers get all users from table users.
func (r RepositoryImpl) GetAllUsers() ([]models.User, error) {
	result, err := r.DB.Query("select id, email from users")
	if err != nil {
		return []models.User{}, err
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
	return users, nil
}

func (r RepositoryImpl) GetUser(email string) (models.User, error) {
	var id int64
	var userEmail string
	err := r.DB.QueryRow("select id, email from users where email=$1", email).Scan(&id, &userEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user does not exist")
		} else {
			return models.User{}, err
		}

	}
	user := models.User{ID: id, Email: userEmail}
	return user, nil
}
