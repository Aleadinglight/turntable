package utils

import "strings"

func InsertLineBreaks(s string, interval int) string {
	var result strings.Builder
	for i, r := range s {
		if i%interval == 0 && i != 0 {
			result.WriteRune('\n')
		}
		result.WriteRune(r)
	}
	return result.String()
}
