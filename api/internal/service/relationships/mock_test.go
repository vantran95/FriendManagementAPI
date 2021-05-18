package relationships

import (
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type (
	mockUserServiceRetriever struct {
		TestF        *testing.T
		GetUserInput struct {
			Group []struct {
				Input  string
				Output *models.User
				Err    error
			}
		}
	}

	mockCreateRepository struct {
		TestF          *testing.T
		CreateRelInput struct {
			Input  models.Relationship
			Output bool
			Err    error
		}
	}

	mockRetrieveRepository struct {
		TestF                *testing.T
		GetRelationshipInput struct {
			RequestInput int64
			TargetInput  int64
			Output       *[]models.Relationship
			Err          error
		}
		GetFriendsListInput struct {
			Group []struct {
				Input  int64
				Output *[]models.User
				Err    error
			}
		}
	}
)

func (m mockUserServiceRetriever) GetUser(email string) (*models.User, error) {
	for _, m := range m.GetUserInput.Group {
		if email == m.Input {
			return m.Output, m.Err
		}
	}

	return nil, nil
}

func (m mockCreateRepository) CreateRelationship(relationship models.Relationship) (bool, error) {
	assert.Equal(m.TestF, m.CreateRelInput.Input, relationship)

	return m.CreateRelInput.Output, m.CreateRelInput.Err
}

func (m mockRetrieveRepository) GetRelationships(requestID, targetID int64) (*[]models.Relationship, error) {
	assert.Equal(m.TestF, m.GetRelationshipInput.RequestInput, requestID)
	assert.Equal(m.TestF, m.GetRelationshipInput.TargetInput, targetID)

	return m.GetRelationshipInput.Output, m.GetRelationshipInput.Err
}

func (m mockRetrieveRepository) GetFriendsList(emailID int64) (*[]models.User, error) {
	for _, m := range m.GetFriendsListInput.Group {
		if emailID == m.Input {
			return m.Output, m.Err
		}
	}

	return nil, nil
}
