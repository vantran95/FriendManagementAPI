package relationship

import (
	"InternalUserManagement/models"
	"InternalUserManagement/pkg/enum"
)

// CreateRelationship attempts to retrieve a friend relationship between 2 emails
func (r RelationshipServiceImpl) CreateRelationship(relationship models.Relationship) bool {
	return r.RelationshipRepository.CreateRelationship(relationship)
}

// FindByTwoEmailIdsAndStatus attempts to retrieve a list friend relationship by two emails and status
func (r RelationshipServiceImpl) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error) {
	return r.RelationshipRepository.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, status)
}

// FindByEmailIdAndStatus attempts to retrieve a friend relationship by email address and status
func (r RelationshipServiceImpl) FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error) {
	return r.RelationshipRepository.FindByEmailIdAndStatus(emailId, status)
}

// IsFriended check two emails is friend
func (r RelationshipServiceImpl) IsFriended(firstEmailId int64, secondEmailId int64) bool {
	relationships, _ := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.FRIEND})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}
