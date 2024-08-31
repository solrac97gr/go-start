package utils_test

import (
	"testing"

	"github.com/solrac97gr/go-start/utils"
)

func Test_Capitalize(t *testing.T) {
	var expected = "Test"
	var result = utils.Capitalize("test")
	if result != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, result)
	}
}

func Test_LowerCase(t *testing.T) {
	var expected = "test"
	var result = utils.Lowercase("Test")
	if result != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, result)
	}
}
