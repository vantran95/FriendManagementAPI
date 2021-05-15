package v1

import (
	"fmt"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type mockUserRetrieverService struct {
	TestF            *testing.T
	GetAllUsersInput struct {
		Output []models.User
		Err    error
	}
}

type mockUserCreatorService struct {
	TestF           *testing.T
	CreateUserInput struct {
		Input  string
		Output bool
		Err    error
	}
}

type mockRelationshipCreatorSrv struct {
	TestF           *testing.T
	MakeFriendInput struct {
		FirstInput  string
		SecondInput string
		Output      bool
		Err         error
	}
}

type mockRelationshipRetrieveSrv struct {
	TestF               *testing.T
	GetFriendsListInput struct {
		Input  string
		Output []string
		Err    error
	}
	GetCommonFriendsInput struct {
		FirstInput  string
		SecondInput string
		Output      []string
		Err         error
	}
}

func (m mockUserRetrieverService) GetAllUsers() ([]models.User, error) {
	fmt.Println("fired to mock CreateUser")
	fmt.Println("mock result: ", m.GetAllUsersInput.Output)
	return m.GetAllUsersInput.Output, m.GetAllUsersInput.Err
}

func (m mockUserRetrieverService) GetUser(string) (models.User, error) {
	return models.User{}, nil
}

func (m mockUserCreatorService) CreateUser(email string) (bool, error) {
	assert.Equal(m.TestF, m.CreateUserInput.Input, email)

	return m.CreateUserInput.Output, m.CreateUserInput.Err
}

func (m mockRelationshipCreatorSrv) MakeFriend(firstEmail, secondEmail string) (bool, error) {
	fmt.Println("fired to mock MakeFriend")
	fmt.Println("mock result: ", m.MakeFriendInput.Output)

	assert.Equal(m.TestF, m.MakeFriendInput.FirstInput, firstEmail)
	assert.Equal(m.TestF, m.MakeFriendInput.SecondInput, secondEmail)

	return m.MakeFriendInput.Output, m.MakeFriendInput.Err
}

func (m mockRelationshipRetrieveSrv) GetFriendsList(email string) ([]string, error) {
	fmt.Println("fired to mock MakeFriend")
	fmt.Println("mock result: ", m.GetFriendsListInput.Output)

	assert.Equal(m.TestF, m.GetFriendsListInput.Input, email)
	return m.GetFriendsListInput.Output, m.GetFriendsListInput.Err
}

func (m mockRelationshipRetrieveSrv) GetCommonFriends(firstEmail, secondEmail string) ([]string, error) {
	fmt.Println("fired to mock MakeFriend")
	fmt.Println("mock result: ", m.GetCommonFriendsInput.Output)

	assert.Equal(m.TestF, m.GetCommonFriendsInput.FirstInput, firstEmail)
	assert.Equal(m.TestF, m.GetCommonFriendsInput.SecondInput, secondEmail)
	return m.GetCommonFriendsInput.Output, m.GetCommonFriendsInput.Err
}
