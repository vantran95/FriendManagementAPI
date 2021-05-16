package relationships

import (
	"fmt"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockUserServiceRetriever struct {
	TestF        *testing.T
	GetUserInput struct {
		Input  string
		Output *models.User
		Err    error
	}
}

type mockCreateRepository struct {
	TestF                *testing.T
	GetRelationshipInput struct {
		FromInput int64
		ToInput   int64
		Output    *[]models.Relationship
		Err       error
	}
	CreateRelInput struct {
		Input  models.Relationship
		Output bool
		Err    error
	}
}

type mockRetrieveRepository struct {
	TestF                *testing.T
	GetRelationshipInput struct {
		FromInput int64
		ToInput   int64
		Output    *[]models.Relationship
		Err       error
	}
	GetFriendsListInput struct {
		Input  int64
		Output *[]models.User
		Err    error
	}
}

func (m mockUserServiceRetriever) GetUser(email string) (*models.User, error) {
	assert.Equal(m.TestF, m.GetUserInput.Input, email)

	return m.GetUserInput.Output, m.GetUserInput.Err
}

func (m mockCreateRepository) CreateRelationship(relationship models.Relationship) (bool, error) {
	assert.Equal(m.TestF, m.CreateRelInput.Input, relationship)

	return m.CreateRelInput.Output, m.CreateRelInput.Err
}
func (m mockCreateRepository) GetRelationships(fromID, toID int64) (*[]models.Relationship, error) {
	assert.Equal(m.TestF, m.GetRelationshipInput.FromInput, fromID)
	assert.Equal(m.TestF, m.GetRelationshipInput.ToInput, toID)

	return m.GetRelationshipInput.Output, m.GetRelationshipInput.Err
}

func (m mockRetrieveRepository) GetRelationships(fromID, toID int64) (*[]models.Relationship, error) {
	assert.Equal(m.TestF, m.GetRelationshipInput.FromInput, fromID)
	assert.Equal(m.TestF, m.GetRelationshipInput.ToInput, toID)

	return m.GetRelationshipInput.Output, m.GetRelationshipInput.Err
}

func (m mockRetrieveRepository) GetFriendsList(emailID int64) (*[]models.User, error) {
	fmt.Println("fired to mock MakeFriend")
	fmt.Println("mock result: ", m.GetFriendsListInput.Output)

	assert.Equal(m.TestF, m.GetFriendsListInput.Input, emailID)
	return m.GetFriendsListInput.Output, m.GetFriendsListInput.Err
}
