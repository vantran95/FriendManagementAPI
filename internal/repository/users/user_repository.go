package users

//import "database/sql"

// NewUserRepository use for testing
//func NewUserRepository(db *sql.DB) UserRepo  {
//	return UserRepo{DB: db}
//}

func (u UserRepositoryImpl) GetAllUsers() []string {
	result, err := u.DB.Query("select email from user_management")
	if err != nil {
		panic(err.Error())
	}

	var emails []string

	for result.Next() {
		var email string
		err = result.Scan(email)
		if err != nil {
			panic(err.Error())
		}
		emails = append(emails, email)
	}
	return emails
}
