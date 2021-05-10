package utils

import "regexp"

const EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func IsFormatEmail(email string) bool {
	re, _ := regexp.Compile(EMAIL_REGEX)
	if re.MatchString(email) {
		return true
	}
	return false
}
