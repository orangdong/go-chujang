package models

type User struct {
	ID       int     `db:"id"`
	Name     string  `db:"name"`
	Username string  `db:"username"`
	Password string  `db:"password"`
	Email    string  `db:"email"`
	Avatar   *string `db:"avatar"`
}

type UserDTO struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Avatar   *string `json:"avatar"`
}

type UserCreate struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
}
