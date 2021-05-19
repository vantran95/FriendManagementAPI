package users

import (
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type (
	// mockCreateRepository stores info to mock the create repository
	mockCreateRepository struct {
		TestF        *testing.T
		GetUserInput struct {
			Input  string
			Output *models.User
			Err    error
		}
		CreateUserInput struct {
			Input  string
			Output bool
			Err    error
		}
	}

	// mockCreateRepository stores info to mock the retrieve repository
	mockRetrieveRepository struct {
		TestF        *testing.T
		GetUserInput struct {
			Input  string
			Output *models.User
			Err    error
		}
		GetAllUsersInput struct {
			Output *[]models.User
			Err    error
		}
	}
)

// GetUser attempts to mock the get user from create repository
func (m mockCreateRepository) GetUser(email string) (*models.User, error) {
	assert.Equal(m.TestF, m.GetUserInput.Input, email)
	return m.GetUserInput.Output, m.GetUserInput.Err
}

// CreateUser attempts to mock the create user from create repository
func (m mockCreateRepository) CreateUser(email string) (bool, error) {
	assert.Equal(m.TestF, m.CreateUserInput.Input, email)
	return m.CreateUserInput.Output, m.CreateUserInput.Err
}

// GetUser attempts to mock the get user from the retrieve repository
func (m mockRetrieveRepository) GetUser(email string) (*models.User, error) {
	assert.Equal(m.TestF, m.GetUserInput.Input, email)
	return m.GetUserInput.Output, m.GetUserInput.Err
}

// GetAllUsers attempts to mock the get all users from the retrieve repository
func (m mockRetrieveRepository) GetAllUsers() (*[]models.User, error) {
	return m.GetAllUsersInput.Output, m.GetAllUsersInput.Err
}
