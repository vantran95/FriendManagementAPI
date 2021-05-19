package v1

import (
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type (
	// mockUserRetrieverService stores info to mock the retriever user service
	mockUserRetrieverService struct {
		TestF            *testing.T
		GetAllUsersInput struct {
			Output []models.User
			Err    error
		}
	}
	// mockUserCreatorService stores info to mock the create user service
	mockUserCreatorService struct {
		TestF           *testing.T
		CreateUserInput struct {
			Input  string
			Output bool
			Err    error
		}
	}

	// mockRelationshipCreatorSrv stores info to mock the create relationship service
	mockRelationshipCreatorSrv struct {
		TestF           *testing.T
		MakeFriendInput struct {
			RequestInput string
			TargetInput  string
			Output       bool
			Err          error
		}
	}

	// mockRelationshipRetrieveSrv stores info to mock the retriever relationship service
	mockRelationshipRetrieveSrv struct {
		TestF               *testing.T
		GetFriendsListInput struct {
			Input  string
			Output []string
			Err    error
		}
		GetCommonFriendsInput struct {
			RequestInput string
			TargetInput  string
			Output       []string
			Err          error
		}
	}
)

// GetAllUsers attempts to mock get all users from the retriever user service
func (m mockUserRetrieverService) GetAllUsers() ([]models.User, error) {
	return m.GetAllUsersInput.Output, m.GetAllUsersInput.Err
}

// CreateUser attempts to mock create user from the create user service
func (m mockUserCreatorService) CreateUser(email string) (bool, error) {
	assert.Equal(m.TestF, m.CreateUserInput.Input, email)
	return m.CreateUserInput.Output, m.CreateUserInput.Err
}

// MakeFriend attempts to mock make friend from the create relationship service
func (m mockRelationshipCreatorSrv) MakeFriend(requestEmail, targetEmail string) (bool, error) {
	assert.Equal(m.TestF, m.MakeFriendInput.RequestInput, requestEmail)
	assert.Equal(m.TestF, m.MakeFriendInput.TargetInput, targetEmail)
	return m.MakeFriendInput.Output, m.MakeFriendInput.Err
}

// GetFriendsList attempts to mock get friends list from the retriever relationship service
func (m mockRelationshipRetrieveSrv) GetFriendsList(email string) ([]string, error) {
	assert.Equal(m.TestF, m.GetFriendsListInput.Input, email)
	return m.GetFriendsListInput.Output, m.GetFriendsListInput.Err
}

// GetCommonFriends attempts to mock to get common friends from the retriever relationship service
func (m mockRelationshipRetrieveSrv) GetCommonFriends(requestEmail, targetEmail string) ([]string, error) {
	assert.Equal(m.TestF, m.GetCommonFriendsInput.RequestInput, requestEmail)
	assert.Equal(m.TestF, m.GetCommonFriendsInput.TargetInput, targetEmail)
	return m.GetCommonFriendsInput.Output, m.GetCommonFriendsInput.Err
}
