package users

type Repository interface {
	GetAllUsers() []string
}

type ServiceImpl struct {
	Repository Repository
}
