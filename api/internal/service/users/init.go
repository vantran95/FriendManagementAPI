package users

// ServiceImpl stores info to retrieve user service.
type ServiceImpl struct {
	CreateRepo   createRepository
	RetrieveRepo retrieveRepository
}
