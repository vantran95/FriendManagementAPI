package relationships

// ServiceImpl stores info to retrieve relationship service.
type ServiceImpl struct {
	CreateRepo        createRepository
	RetrieveRepo      retrieveRepository
	UserRetrieverRepo userRetrieverRepo
}
