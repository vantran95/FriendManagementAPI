package users

type Repository interface {
	GetAllUsers() []string
	CreateUser(email string) bool
	ExistsByEmail(email string) bool
}

type ServiceImpl struct {
	Repository Repository
}
