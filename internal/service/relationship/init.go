package relationship

import (
	"InternalUserManagement/models"
)

type RelationshipRepository interface {
	CreateRelationship(relationship models.Relationship) bool
}

type RelationshipServiceImpl struct {
	RelationshipRepository RelationshipRepository
}
