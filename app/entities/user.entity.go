package entities

type UserModel struct {
	ID        string  `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	Username  string  `db:"username" json:"username"`
	Email     string  `db:"email" json:"email"`
	Avatar    *string `db:"avatar" json:"avatar"`
	CreatedAt *string `db:"created_at" json:"createdAt"`
	UpdatedAt *string `db:"updated_at" json:"updatedAt"`
}

type UserCreate struct {
	Name     string  `json:"name" validate:"required"`
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Avatar   *string `json:"avatar" validate:"omitempty,url"`
}

type UserUpdate struct {
	Name     *string `json:"name" validate:"omitempty"`
	Username *string `json:"username" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Avatar   *string `json:"avatar" validate:"omitempty,url"`
}
