package users

import "database/sql"

// UserRepositoryImpl stores info to retrieve user repository.
type UserRepositoryImpl struct {
	DB *sql.DB
}
