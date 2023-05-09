package models

// Post without comments
// (seen on link aggregation pages)
type ShortPost struct {
	ShortID          string `json:"short_id"`
	ShortIDURL       string `json:"short_id_url"`
	CreatedAt        string `json:"created_at"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	Score            int    `json:"score"`
	Flags            int    `json:"flags"`
	CommentCount     int    `json:"comment_count"`
	Description      string `json:"description"`
	DescriptionPlain string `json:"description_plain"`
	CommentsURL      string `json:"comments_url"`
	User             `json:"submitter_user"`
	Tags             []string `json:"tags"`
}

// Post including comments
// (seen when you click "comments" on the post)
type Post struct {
	ShortPost
	Comments []Comment `json:"comments"`
}
