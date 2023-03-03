package models

type User struct {
	ID       int64   `json:"id" db:"id"`
	UserName string  `json:"username" db:"username"`
	Password string  `json:"password" db:"password"`
	Email    *string `json:"email" db:"email"`
	Role     int64   `json:"role" db:"role"`
	*Time
}
