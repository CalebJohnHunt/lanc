package dto

import "time"

type Comment struct {
	ShortID        string    `json:"short_id"`
	ShortIDURL     string    `json:"short_id_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
	IsModerated    bool      `json:"is_moderated"`
	Score          int       `json:"score"`
	Flags          int       `json:"flags"`
	ParentComment  *string   `json:"parent_comment"`
	Comment        string    `json:"comment"`
	CommentPlain   string    `json:"comment_plain"`
	URL            string    `json:"url"`
	IndentLevel    int       `json:"indent_level"`
	CommentingUser User      `json:"commenting_user"`
}
