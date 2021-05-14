package v1

import (
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
