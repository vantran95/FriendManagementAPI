package users

import (
	"database/sql"
)

// RepositoryImpl stores info to retrieve user repository.
type RepositoryImpl struct {
	DB *sql.DB
}
