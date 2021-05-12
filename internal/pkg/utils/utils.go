package utils

import (
	"net"
	"regexp"
	"strings"
)

const EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+\\\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// IsFormatEmail attempts to check email format.
func IsFormatEmail(email string) bool {
	re, _ := regexp.Compile(EMAIL_REGEX)

	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !re.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

// RemoveItemFromList attempts to remove items from list.
func RemoveItemFromList(list []int64, item int64) []int64 {
	newList := []int64{}
	for _, i := range list {
		if i != item {
			newList = append(newList, i)
		}
	}
	return newList
}
