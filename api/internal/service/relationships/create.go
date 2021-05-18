package relationships

import (
	"database/sql"
	"errors"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// createRepository interface represents the create a relationship repository
type createRepository interface {
	CreateRelationship(relationship models.Relationship) (bool, error)
}

// MakeFriend attempts to create a friend relationship between two emails.
func (s ServiceImpl) MakeFriend(requestEmail, targetEmail string) (bool, error) {
	requestUser, err := s.UserRetriever.GetUser(requestEmail)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return false, errors.New("user does not exists")
		case err != nil:
			return false, err
		}
	}

	targetUser, err := s.UserRetriever.GetUser(targetEmail)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return false, errors.New("user does not exists")
		case err != nil:
			return false, err
		}
	}

	// Get relationship and check friend
	rs, err := getRelationships(s.RetrieveRepo, requestUser.ID, targetUser.ID)
	if err != nil {
		return false, err
	}

	for _, item := range *rs {
		switch item.Status {
		case RelationshipTypeFriend:
			return false, errors.New("already friended")
		case RelationshipTypeBlocked:
			return false, errors.New("you were blocked")
		}
	}

	relationship := models.Relationship{RequestID: requestUser.ID, TargetID: targetUser.ID, Status: RelationshipTypeFriend}

	return s.CreateRepo.CreateRelationship(relationship)
}
