package relationship

import (
	"InternalUserManagement/models"
)

// Repository interface represents the criteria used to retrieve a relationship repository.
type Repository interface {
	CreateRelationship(relationship models.Relationship) (bool, error)
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error)
	FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error)
}

// UserService interface represents the criteria used to retrieve a user service.
type UserService interface {
	ExistsByEmail(email string) (bool, error)
	FindUserIdByEmail(email string) (int64, error)
	FindEmailByIds(ids []int64) ([]string, error)
}

// ServiceImpl stores info to retrieve relationship service.
type ServiceImpl struct {
	Repository  Repository
	UserService UserService
}
