package util

import (
	"strings"
)

func Contains(s []string, substr string) bool {
	for _, v := range s {
		if v == substr {
			return true
		}
	}
	return false
}

func ContainsStartWith(s []string, subStr string) bool {
	for _, v := range s {
		if strings.HasPrefix(subStr, v) {
			return true
		}
	}
	return false
}

func ContainsEndWith(s []string, subStr string) bool {
	for _, v := range s {
		if strings.HasSuffix(subStr, v) {
			return true
		}
	}
	return false
}
