package relationship

import (
	"InternalUserManagement/models"
)

type RelationshipRepository interface {
	CreateRelationship(relationship models.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error)
	FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error)
}

type RelationshipServiceImpl struct {
	RelationshipRepository RelationshipRepository
}
