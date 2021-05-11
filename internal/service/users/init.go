package users

type Repository interface {
	GetAllUsers() []string
	CreateUser(email string) bool
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) int64
}

type ServiceImpl struct {
	Repository Repository
}
