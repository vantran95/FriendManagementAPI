package relationship

import "database/sql"

type RelationshipRepositoryImpl struct {
	DB *sql.DB
}
