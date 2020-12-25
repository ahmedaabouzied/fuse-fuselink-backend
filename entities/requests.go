package entities

const (
	UserIDContextKey string = "userID"
)

type CreateUserRequest struct {
	Username string `json:"username"`
}

type UpdateUserRequest struct {
	Username       string          `json:"username"`
	SocialAccounts []SocialAccount `json:"social_accounts"`
}
