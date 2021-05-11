package relationship

import (
	"InternalUserManagement/models"
)

// RelationshipRepository interface represents the criteria used to retrieve a relationship repository.
type RelationshipRepository interface {
	CreateRelationship(relationship models.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error)
	FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error)
}

// RelationshipServiceImpl stores info to retrieve relationship service.
type RelationshipServiceImpl struct {
	RelationshipRepository RelationshipRepository
}
