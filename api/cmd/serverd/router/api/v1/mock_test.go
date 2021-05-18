package v1

import (
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type (
	mockUserRetrieverService struct {
		TestF            *testing.T
		GetAllUsersInput struct {
			Output []models.User
			Err    error
		}
	}

	mockUserCreatorService struct {
		TestF           *testing.T
		CreateUserInput struct {
			Input  string
			Output bool
			Err    error
		}
	}

	mockRelationshipCreatorSrv struct {
		TestF           *testing.T
		MakeFriendInput struct {
			RequestInput string
			TargetInput  string
			Output       bool
			Err          error
		}
	}

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

func (m mockUserRetrieverService) GetAllUsers() ([]models.User, error) {
	return m.GetAllUsersInput.Output, m.GetAllUsersInput.Err
}

func (m mockUserRetrieverService) GetUser(string) (models.User, error) {
	return models.User{}, nil
}

func (m mockUserCreatorService) CreateUser(email string) (bool, error) {
	assert.Equal(m.TestF, m.CreateUserInput.Input, email)

	return m.CreateUserInput.Output, m.CreateUserInput.Err
}

func (m mockRelationshipCreatorSrv) MakeFriend(requestEmail, targetEmail string) (bool, error) {
	assert.Equal(m.TestF, m.MakeFriendInput.RequestInput, requestEmail)
	assert.Equal(m.TestF, m.MakeFriendInput.TargetInput, targetEmail)

	return m.MakeFriendInput.Output, m.MakeFriendInput.Err
}

func (m mockRelationshipRetrieveSrv) GetFriendsList(email string) ([]string, error) {
	assert.Equal(m.TestF, m.GetFriendsListInput.Input, email)

	return m.GetFriendsListInput.Output, m.GetFriendsListInput.Err
}

func (m mockRelationshipRetrieveSrv) GetCommonFriends(requestEmail, targetEmail string) ([]string, error) {
	assert.Equal(m.TestF, m.GetCommonFriendsInput.RequestInput, requestEmail)
	assert.Equal(m.TestF, m.GetCommonFriendsInput.TargetInput, targetEmail)

	return m.GetCommonFriendsInput.Output, m.GetCommonFriendsInput.Err
}
