package utils

import (
	"strings"
)

// Capitalize returns the string with the first letter in uppercase
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

// Lowercase returns the string in Lowercase
func Lowercase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s)
}
