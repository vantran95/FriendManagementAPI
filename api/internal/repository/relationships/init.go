package relationships

import "database/sql"

// RepositoryImpl stores info to retrieve relationship repository.
type RepositoryImpl struct {
	DB *sql.DB
}
