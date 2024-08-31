package utils

import (
	"strings"
)

// Capitalize returns the string with the first letter in uppercase
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	if containsDash(s) {
		return ProcessCases(s)
	}
	return strings.Title(strings.ToLower(s))
}

// Lowercase returns the string in Lowercase
func Lowercase(s string) string {
	if len(s) == 0 {
		return s
	}
	if containsDash(s) {
		return ProcessCasesToLowercase(s)
	}
	return strings.ToLower(s)
}

// Process Cases new-package -> NewPackage
func ProcessCases(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

// ProcessCasesToLowercase new-package -> newPackage
func ProcessCasesToLowercase(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	return strings.ToLower(s)
}

// Contains a "-"
func containsDash(s string) bool {
	if strings.Contains(s, "-") {
		return true
	}
	return false
}
