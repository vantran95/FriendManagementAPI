package models

// User stores info to retrieve user struct
type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
