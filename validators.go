package w

import (
	"strings"
)

// IsBlank returns true if the strings is zero length or if the string contains
// only whitespace characters.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsLengthAtLeast returns true if length of s is at least length characters
// long.
func IsLengthAtLeast(length int, s string) bool {
	return len(s) >= length
}

// IsLengthAtMost returns true if length of s is at most length characters long.
func IsLengthAtMost(length int, s string) bool {
	return len(s) <= length
}
