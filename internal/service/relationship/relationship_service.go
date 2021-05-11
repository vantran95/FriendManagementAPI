package relationship

import (
	"InternalUserManagement/models"
)

func (r RelationshipServiceImpl) CreateRelationship(relationship models.Relationship) bool {
	return r.RelationshipRepository.CreateRelationship(relationship)
}
