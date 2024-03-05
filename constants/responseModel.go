package constants

type Response struct {
	Status string `json:"status"`
	Error  string `json:"debug_errors"`
	Data   any    `json:"data"`
}

type LoginResponse struct {
	UserID      uint   `json:"user_id"`
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
}

type RegisterResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}
