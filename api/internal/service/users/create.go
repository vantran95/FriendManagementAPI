package users

import "errors"

// CreateUser attempts to create a user
func (s ServiceImpl) CreateUser(email string) (bool, error) {
	getUser, err := s.Repository.GetUser(email)
	if err != nil {
		return false, err
	}

	if getUser.Email != "" {
		return false, errors.New("email already exists")
	}
	return s.Repository.CreateUser(email)
}
