package relationships

import "github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"

// Repository interface represents the criteria used to retrieve a relationship repository.
type Repository interface {
	CreateRelationship(relationship models.Relationship) (bool, error)
	GetRelationships(fromID, toID int64) ([]models.Relationship, error)
	GetFriendsList(emailID int64) ([]models.User, error)
}

// UserService interface represents the criteria used to retrieve a user service.
type UserService interface {
	GetUser(email string) (models.User, error)
}

// ServiceImpl stores info to retrieve relationship service.
type ServiceImpl struct {
	Repository  Repository
	UserService UserService
}
