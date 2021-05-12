package utils

import (
	"regexp"
)

const EMAIL_REGEX = "[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})"

// IsFormatEmail attempts to check email format.
func IsFormatEmail(email string) bool {
	re, _ := regexp.Compile(EMAIL_REGEX)

	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !re.MatchString(email) {
		return false
	}
	return true
}
