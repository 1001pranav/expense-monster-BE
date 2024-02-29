package constants

type LoginResponse struct {
	UserID      uint   `json:"user_id"`
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
}
