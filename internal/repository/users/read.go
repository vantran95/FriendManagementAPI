package users

import "InternalUserManagement/models"

// GetAllUsers get all users from table users.
func (r RepositoryImpl) GetAllUsers() ([]models.User, error) {
	result, err := r.DB.Query("select * from users")
	if err != nil {
		return []models.User{}, err
	}

	var users []models.User

	for result.Next() {
		var id int64
		var email string
		result.Scan(&id, &email)
		user := models.User{ID: id, Email: email}
		users = append(users, user)
	}
	return users, nil
}

func (r RepositoryImpl) GetUser(email string) (models.User, error) {
	var id int64
	var userEmail string
	err := r.DB.QueryRow("select * from users where email = $1", email).Scan(&id, &userEmail)

	if err != nil {
		return models.User{}, err
	}
	user := models.User{ID: id, Email: userEmail}
	return user, nil
}
