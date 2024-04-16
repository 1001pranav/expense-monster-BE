package constants

type APIRequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type APIRequestForgotPassword struct {
	Email string `json:"email"`
}

type APIRequestResetPassword struct {
	OTP          *string          `json:"otp"`
	Password     string           `json:"password"`
	OldPassword  *string          `json:"old_password"`
	PasswordType RestPasswordType `json:"password_type"`
	Email        string           `json:"email"`
}
