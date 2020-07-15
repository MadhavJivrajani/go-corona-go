package utils

import (
	"strings"
)

// ExistsString checks if a string is present in a slice of strings.
func ExistsString(items []string, key string) (string, bool) {
	for _, value := range items {
		temp := strings.ToLower(value)
		if temp == key {
			return value, true
		}
	}
	return "", false
}
