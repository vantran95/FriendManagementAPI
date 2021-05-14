package users

import (
	"testing"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type mockCreateRepository struct {
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

func (m mockCreateRepository) GetUser(email string) (*models.User, error) {
	assert.Equal(m.TestF, m.GetUserInput.Input, email)

	return m.GetUserInput.Output, m.GetUserInput.Err
}

func (m mockCreateRepository) CreateUser(email string) (bool, error) {
	assert.Equal(m.TestF, m.CreateUserInput.Input, email)

	return m.CreateUserInput.Output, m.CreateUserInput.Err
}
