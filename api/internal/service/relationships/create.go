package relationships

import (
	"errors"
	"fmt"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

// createRepository interface represents the create a relationship repository
type createRepository interface {
	CreateRelationship(relationship models.Relationship) (bool, error)
}

// MakeFriend attempts to create a friend relationship between two emails.
func (s ServiceImpl) MakeFriend(firstEmail, secondEmail string) (bool, error) {
	firstUser, err := s.UserRetriever.GetUser(firstEmail)
	if err != nil {
		return false, err
	}

	secondUser, err := s.UserRetriever.GetUser(secondEmail)
	if err != nil {
		return false, err
	}

	// Get relationship and check friend
	rs, err := getRelationships(s.RetrieveRepo, firstUser.ID, secondUser.ID)
	if err != nil {
		return false, err
	}

	fmt.Println("passed to getRelationships")
	if rs != nil {
		for _, item := range *rs {
			switch item.Status {
			case RelationshipTypeFriend:
				return false, errors.New("already friended")
			case RelationshipTypeBlocked:
				return false, errors.New("you were blocked")
			}
		}
	}

	relationship := models.Relationship{FirstEmailID: firstUser.ID, SecondEmailID: secondUser.ID, Status: RelationshipTypeFriend}

	return s.CreateRepo.CreateRelationship(relationship)
}
