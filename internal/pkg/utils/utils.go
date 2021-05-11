package utils

import "regexp"

const EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// IsFormatEmail attempts to check email format.
func IsFormatEmail(email string) bool {
	re, _ := regexp.Compile(EMAIL_REGEX)
	if re.MatchString(email) {
		return true
	}
	return false
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
