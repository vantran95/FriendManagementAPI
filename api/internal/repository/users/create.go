package users

// CreateUser create email into table users.
func (r RepositoryImpl) CreateUser(email string) (bool, error) {
	query, err := r.DB.Prepare(`insert into users (email) values ($1)`)

	if err != nil {
		return false, err
	}

	_, err = query.Exec(email)

	return err == nil, err
}
