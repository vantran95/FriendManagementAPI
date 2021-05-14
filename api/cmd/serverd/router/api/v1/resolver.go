package v1

type RetrieveResolver struct {
	UserService         userRetrieverService
	RelationshipService relationshipRetrieverService
}

type CreateResolver struct {
	UserService         userCreatorService
	RelationshipService relationshipCreatorService
}
