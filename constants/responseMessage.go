package constants

import "strings"

const (
	DATA_INVALID_RAW     string = "<replace> is invalid, Please fix it and try again"
	REQUIRED_DATA_RAW    string = "Need more information to complete the request, Required <replace> parameter"
	DATA_PRESENT_RAW     string = "<replace> already exists, Please check again"
	DATA_NOT_PRESENT_RAW string = "Sorry 😭, But <replace> does not exist"
)

const (
	INVALID_REQUEST_STATUS        string = "INVALID_REQUEST"
	STATUS_INVALID_EMAIL          string = "INVALID_EMAIL_ID"
	REQUIRED_EMAIL_STATUS         string = "REQUIRED_EMAIL"
	REQUIRED_PASSWORD_STATUS      string = "REQUIRED_PASSWORD"
	INVALID_EMAIL_PASSWORD_STATUS string = "INVALID_EMAIL_OR_PASSWORD"

	EMAIL_EXISTS_STATUS    string = "EMAIL_EXISTS"
	USER_NOT_EXISTS_STATUS string = "USER_NOT_EXISTS"

	STATUS_MISSING_PASSWORD_TYPE string = "MISSING_PASSWORD_TYPE"
	STATUS_MISSING_OTP           string = "MISSING_OTP"
	STATUS_MISSING_OLD_PASSWORD  string = "MISSING_OLD_PASSWORD"
	STATUS_INVALID_PASSWORD_TYPE string = "INVALID_PASSWORD_TYPE"
	STATUS_INVALID_OTP           string = "INVALID_OTP"
	STATUS_OTP_EXPIRED           string = "OTP_EXPIRED"
	STATUS_PASSWORD_NOT_MATCH    string = "PASSWORD_NOT_MATCHING"

	INTERNAL_SERVER_STATUS string = "INTERNAL_SERVER_ERROR"
	SUCCESS_STATUS         string = "SUCCESS"
)

var (
	BAD_REQUEST_MESSAGE            string = REPLACE_STRINGS(DATA_INVALID_RAW, "Request", "")
	REQUIRED_EMAIL_MESSAGE         string = REPLACE_STRINGS(REQUIRED_DATA_RAW, "EMAIL", "")
	REQUIRED_PASSWORD_MESSAGE      string = REPLACE_STRINGS(REQUIRED_DATA_RAW, "PASSWORD", "")
	INVALID_EMAIL_PASSWORD_MESSAGE string = REPLACE_STRINGS(DATA_INVALID_RAW, "Email or Password", "")
	EMAIL_EXISTS_MESSAGE           string = REPLACE_STRINGS(DATA_PRESENT_RAW, "Email Id entered", "")
	USER_NOT_EXISTS_MESSAGE        string = REPLACE_STRINGS(DATA_NOT_PRESENT_RAW, "Email Id entered", "")
)

func REPLACE_STRINGS(baseString string, replacedString string, replaceTag string) string {
	if replaceTag == "" {
		replaceTag = "<replace>"
	}
	return strings.ReplaceAll(baseString, replaceTag, replacedString)
}
