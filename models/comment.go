package models

type Comment struct {
	ShortID       string `json:"short_id"`
	ShortIDURL    string `json:"short_id_url"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	IsDeleted     string `json:"is_deleted"`
	IsModerated   string `json:"is_moderated"`
	Score         string `json:"score"`
	Flags         string `json:"flags"`
	ParentComment string `json:"parent_comment"`
	Comment       string `json:"comment"`
	CommentPlain  string `json:"comment_plain"`
	URL           string `json:"url"`
	IndentLevel   int    `json:"indent_level"`
	User          `json:"commenting_user"`
}
