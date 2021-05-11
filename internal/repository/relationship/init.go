package relationship

import "database/sql"

// RelationshipRepositoryImpl stores info to retrieve relationship repository.
type RelationshipRepositoryImpl struct {
	DB *sql.DB
}
