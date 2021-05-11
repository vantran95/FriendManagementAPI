package relationship

import "InternalUserManagement/models"

func (r RelationshipRepositoryImpl) CreateRelationship(relationship models.Relationship) bool {
	query, err := r.DB.Prepare("insert into relationship (first_email_id, second_email_id, status) values ($1, $2, $3);")

	if err != nil {
		return false
	}

	query.Exec(relationship.FirstEmailId, relationship.SecondEmailID, relationship.Status)
	return true
}
