package entities

type UserModel struct {
	ID       int     `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Username string  `db:"username" json:"username"`
	Password string  `db:"password" `
	Email    string  `db:"email" json:"email"`
	Avatar   *string `db:"avatar" json:"avatar"`
}

type UserCreate struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
}
