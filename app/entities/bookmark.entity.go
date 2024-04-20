package entities

type BookmarkModel struct {
	ID        string  `db:"id" json:"id"`
	UserID    string  `db:"user_id" json:"user_id"`
	Title     string  `db:"title" json:"title"`
	Summary   *string `db:"summary" json:"summary"`
	URL       string  `db:"url" json:"url"`
	KeyWords  string  `db:"keywords" json:"keywords"`
	IsPrivate *bool   `db:"is_private" json:"is_private"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
