package relationships

import (
	"errors"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// MakeFriend attempts to create a relationship between two emails.
func (s ServiceImpl) MakeFriend(firstEmail, secondEmail string) (bool, error) {
	getFirstUser, err := s.UserService.GetUser(firstEmail)
	if err != nil {
		return false, err
	}

	getSecondUser, err := s.UserService.GetUser(secondEmail)
	if err != nil {
		return false, err
	}

	// Get relationship and check friend
	rs, err := getRelationships(s.Repository, getFirstUser.ID, getSecondUser.ID)
	if err != nil {
		return false, err
	}

	for _, item := range rs {
		switch item.Status {
		case RelationshipTypeFriend:
			return false, errors.New("already friended")
		case RelationshipTypeBlocked:
			return false, errors.New("you were blocked")
		}
	}

	relationship := models.Relationship{FirstEmailID: getFirstUser.ID, SecondEmailID: getSecondUser.ID, Status: RelationshipTypeFriend}

	return s.Repository.CreateRelationship(relationship)
}
