package users

func (u ServiceImpl) GetAllUsers() []string {
	return u.Repository.GetAllUsers()
}
