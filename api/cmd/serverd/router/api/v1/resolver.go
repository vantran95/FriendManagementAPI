package v1

type (
	// RetrieveResolver stores info to retrieve a resolver
	RetrieveResolver struct {
		UserService         userRetrieverService
		RelationshipService relationshipRetrieverService
	}

	// CreateResolver stores info to create a resolver
	CreateResolver struct {
		UserService         userCreatorService
		RelationshipService relationshipCreatorService
	}
)
