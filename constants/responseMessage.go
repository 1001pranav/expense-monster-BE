package constants

import "strings"

const (
	DATA_INVALID  string = "<replace> is invalid, Please fix it and try again"
	REQUIRED_DATA string = "Need more information to complete the request, Required <replace> parameter"
)

const (
	INVALID_REQUEST_STATUS        string = "INVALID_REQUEST"
	MISSING_EMAIL_STATUS          string = "REQUIRED_EMAIL"
	MISSING_PASSWORD_STATUS       string = "REQUIRED_PASSWORD"
	REQUIRED_EMAIL_STATUS         string = "REQUIRED_EMAIL"
	REQUIRED_PASSWORD_STATUS      string = "REQUIRED_PASSWORD"
	INVALID_EMAIL_PASSWORD_STATUS string = "INVALID_EMAIL_OR_PASSWORD"
	INTERNAL_SERVER_STATUS        string = "INTERNAL_SERVER_ERROR"
	SUCCESS_STATUS                string = "SUCCESS"
)

var (
	BAD_REQUEST_MESSAGE            string = REPLACE_STRINGS(DATA_INVALID, "Request")
	REQUIRED_EMAIL_MESSAGE         string = REPLACE_STRINGS(REQUIRED_DATA, "EMAIL")
	REQUIRED_PASSWORD_MESSAGE      string = REPLACE_STRINGS(REQUIRED_DATA, "PASSWORD")
	INVALID_EMAIL_PASSWORD_MESSAGE string = REPLACE_STRINGS(DATA_INVALID, "Email or Password")
)

func REPLACE_STRINGS(s string, replace string) string {
	return strings.ReplaceAll(s, "<replace>", replace)
}
