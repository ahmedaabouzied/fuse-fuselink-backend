package auth

type AuthHandler interface {
	ParseAuthToken(token string) (map[string]interface{}, error)
}
