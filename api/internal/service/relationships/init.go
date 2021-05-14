package relationships

// ServiceImpl stores info to retrieve relationship service.
type ServiceImpl struct {
	CreateRepo           createRepository
	RetrieveRepo         retrieveRepository
	UserServiceRetriever userServiceRetriever
}

// Repository interface represents the criteria used to retrieve a relationship repository.
//type Repository interface {
//	CreateRelationship(relationship models.Relationship) (bool, error)
//	GetRelationships(fromID, toID int64) ([]models.Relationship, error)
//	GetFriendsList(emailID int64) ([]models.User, error)
//}
