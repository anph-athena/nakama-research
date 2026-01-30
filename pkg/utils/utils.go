package utils

import (
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
