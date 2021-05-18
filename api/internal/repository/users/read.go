package users

import (
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// GetAllUsers get all users from table users.
func (r RepositoryImpl) GetAllUsers() (*[]models.User, error) {
	result, err := r.DB.Query("select id, email from users")
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	for result.Next() {
		var id int64
		var email string

		err = result.Scan(&id, &email)
		if err == nil {
			users = append(users, models.User{ID: id, Email: email})
		}
	}
	return &users, nil
}

// GetUser attempts to retrieve a user info
func (r RepositoryImpl) GetUser(email string) (*models.User, error) {
	var id int64
	var userEmail string
	err := r.DB.QueryRow("select id, email from users where email=$1", email).Scan(&id, &userEmail)

	if err != nil {
		return nil, err
	}

	return &models.User{ID: id, Email: userEmail}, nil
}
