package users

import (
	"database/sql"
	"log"
)

//import "database/sql"

func (u UserRepositoryImpl) GetAllUsers() []string {
	result, err := u.DB.Query("select email from user_management")
	if err != nil {
		panic(err.Error())
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
	return emails
}

func (u UserRepositoryImpl) CreateUser(email string) bool {
	query, err := u.DB.Prepare(`insert into user_management (email) values ($1)`)

	if err != nil {
		panic(err.Error())
	}

	query.Exec(email)
	return true
}

// ExistsByEmail check email is exists.
func (u UserRepositoryImpl) ExistsByEmail(email string) bool {
	var id int

	err := u.DB.QueryRow("select id from user_management where email = $1", email).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false
	}
	return true
}

func (u UserRepositoryImpl) FindUserIdByEmail(email string) int64 {
	var id int64
	err := u.DB.QueryRow("select id from user_management where email = $1", email).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}
