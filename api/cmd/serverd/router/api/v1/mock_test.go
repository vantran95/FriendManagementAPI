package v1

import (
	"fmt"
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct {
	TestF                 *testing.T
	mockGetAllUsersResult []models.User
	mockCreateUserInput   string
	mockCreateUserResult  bool
	mockError             error
}

type mockRelationshipService struct {
	TestF                     *testing.T
	mockMakeFriendFirstInput  string
	mockMakeFriendSecondInput string
	mockMakeFriendResult      bool
	mockError                 error
}

func (m mockUserService) GetAllUsers() ([]models.User, error) {
	fmt.Println("fired to mock GetAllUsers")
	return m.mockGetAllUsersResult, m.mockError
}

func (m mockUserService) CreateUser(email string) (bool, error) {
	fmt.Println("fired to mock CreateUser")
	fmt.Println("mock result: ", m.mockCreateUserResult)

	assert.Equal(m.TestF, m.mockCreateUserInput, email)

	return m.mockCreateUserResult, m.mockError
}

func (m mockRelationshipService) MakeFriend(firstEmail, secondEmail string) (bool, error) {
	fmt.Println("fired to mock MakeFriend")
	fmt.Println("mock result: ", m.mockMakeFriendResult)

	assert.Equal(m.TestF, m.mockMakeFriendFirstInput, firstEmail)
	assert.Equal(m.TestF, m.mockMakeFriendSecondInput, secondEmail)

	return m.mockMakeFriendResult, m.mockError
}

func (m mockRelationshipService) GetFriendsList(email string) ([]string, error) {
	panic("implement me")
}

func (m mockRelationshipService) GetCommonFriends(firstEmail, secondEmail string) ([]string, error) {
	panic("implement me")
}

// GetFriendsList(email string) ([]string, error)
// GetCommonFriends(firstEmail, secondEmail string) ([]string, error)
