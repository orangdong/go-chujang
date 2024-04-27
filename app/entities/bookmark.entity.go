package entities

type BookmarkModel struct {
	ID        string  `db:"id" json:"id"`
	UserID    string  `db:"user_id" json:"userId"`
	Title     string  `db:"title" json:"title"`
	Summary   *string `db:"summary" json:"summary"`
	URL       string  `db:"url" json:"url"`
	KeyWords  string  `db:"keywords" json:"keywords"`
	IsPrivate bool    `db:"is_private" json:"isPrivate"`
	CreatedAt *string `db:"created_at" json:"createdAt"`
	UpdatedAt *string `db:"updated_at" json:"updatedAt"`
}

type BookmarkCreate struct {
	Title     string  `json:"title" validate:"required"`
	Summary   *string `json:"summary" validate:"omitempty"`
	URL       string  `json:"url" validate:"required,url"`
	KeyWords  string  `json:"keywords" validate:"required"`
	IsPrivate bool    `db:"is_private" json:"isPrivate" validate:"omitempty"`
}

type BookmarkUpdate struct {
	Title     *string `json:"title" validate:"omitempty"`
	Summary   *string `json:"summary" validate:"omitempty"`
	URL       *string `json:"url" validate:"omitempty,url"`
	KeyWords  *string `json:"keywords" validate:"omitempty"`
	IsPrivate *bool   `db:"is_private" json:"isPrivate" validate:"omitempty"`
}
