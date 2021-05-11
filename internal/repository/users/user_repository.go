package users

import (
	"fmt"
	"strconv"
	"strings"
)

//import "database/sql"

// GetAllUsers get all users from table user_management.
func (u UserRepositoryImpl) GetAllUsers() ([]string, error) {
	result, err := u.DB.Query("select email from user_management")
	if err != nil {
		return []string{}, err
	}

	var emails []string

	for result.Next() {
		var email string
		err = result.Scan(&email)
		if err != nil {
			panic(err.Error())
		}
		emails = append(emails, email)
	}
	return emails, nil
}

// CreateUser create email into table user_management.
func (u UserRepositoryImpl) CreateUser(email string) (bool, error) {
	query, err := u.DB.Prepare(`insert into user_management (email) values ($1)`)

	if err != nil {
		return false, err
	}

	query.Exec(email)
	return true, nil
}

// ExistsByEmail check email is exists.
func (u UserRepositoryImpl) ExistsByEmail(email string) (bool, error) {
	var id int

	err := u.DB.QueryRow("select id from user_management where email = $1", email).Scan(&id)

	if err != nil {
		return false, err
	}
	return true, nil
}

// FindUserIdByEmail find email id by email
func (u UserRepositoryImpl) FindUserIdByEmail(email string) (int64, error) {
	var id int64
	err := u.DB.QueryRow("select id from user_management where email = $1", email).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// FindEmailByIds find email addresses by email ids
func (u UserRepositoryImpl) FindEmailByIds(ids []int64) ([]string, error) {
	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.FormatInt(id, 10)
	}

	qr := `select x.email
			from user x
			where x.id in (%s);
			`
	query := fmt.Sprintf(qr, strings.Join(strIds, ","))
	results, err := u.DB.Query(query)
	if err != nil {
		return []string{}, err
	}

	emails := []string{}

	for results.Next() {
		var email string
		results.Scan(&email)
		emails = append(emails, email)
	}
	return emails, nil
}
