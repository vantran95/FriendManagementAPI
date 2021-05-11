package users

// Repository interface represents the criteria used to retrieve a user repository.
type Repository interface {
	GetAllUsers() ([]string, error)
	CreateUser(email string) (bool, error)
	ExistsByEmail(email string) (bool, error)
	FindUserIdByEmail(email string) (int64, error)
	FindEmailByIds(ids []int64) ([]string, error)
}

// ServiceImpl stores info to retrieve user service.
type ServiceImpl struct {
	Repository Repository
}
