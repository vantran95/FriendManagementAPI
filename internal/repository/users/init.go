package users

import "database/sql"

type UserRepositoryImpl struct {
	DB *sql.DB
}
