package constants

import "time"

type Response struct {
	Status string `json:"status"`
	Error  string `json:"debug_errors"`
}

type LoginResponseData struct {
	UserID      uint   `json:"user_id"`
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
}

type LoginResponse struct {
	Response
	Data *LoginResponseData `json:"data"`
}

type RegisterResponseData struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

type RegisterResponse struct {
	Response
	Data *RegisterResponseData `json:"data"`
}

type ForgotPasswordResponseData struct {
	UserID       uint      `json:"user_id"`
	OTP          uint      `json:"otp"`
	OTPExpiresOn time.Time `json:"otp_expires_on"`
}

type ForgotPasswordResponse struct {
	Response
	Data *ForgotPasswordResponseData `json:"data"`
}
