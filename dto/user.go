package dto

type User struct {
	Username        string  `json:"username"`
	CreatedAt       string  `json:"created_at"`
	IsAdmin         bool    `json:"is_admin"`
	About           string  `json:"about"`
	IsModerator     bool    `json:"is_moderator"`
	Karma           int     `json:"karma"`
	AvatarURL       string  `json:"avatar_url"`
	InvitedByUser   *string `json:"invited_by_user"`
	GithubUsername  string  `json:"github_username"`
	TwitterUsername string  `json:"twitter_username"`
}
