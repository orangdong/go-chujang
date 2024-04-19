package entities

type UserModel struct {
	ID       string  `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Username string  `db:"username" json:"username"`
	Email    string  `db:"email" json:"email"`
	Avatar   *string `db:"avatar" json:"avatar"`
}

type UserCreate struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
}
