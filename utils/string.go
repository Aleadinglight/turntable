package utils

import (
	"strings"
)

func InsertLineBreaks(s string, interval int) string {
	var result strings.Builder
	charCount := 0
	for _, r := range s {
		result.WriteRune(r)
		charCount++
		if charCount%interval == 0 && charCount != len([]rune(s)) {
			result.WriteRune('\n')
		}
	}
	return result.String()
}
