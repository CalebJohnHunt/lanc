package models

type User struct {
	Username        string  `json:"username"`
	CreatedAt       string  `json:"created_at"`
	IsAdmin         string  `json:"is_admin"`
	About           string  `json:"about"`
	IsModerator     string  `json:"is_moderator"`
	Karma           string  `json:"karma"`
	AvatarURL       string  `json:"avatar_url"`
	InvitedByUser   *string `json:"invited_by_user"`
	GithubUsername  string  `json:"github_username"`
	TwitterUsername string  `json:"twitter_username"`
}
