package utils

import "strings"

func TrimSpaces(str *string) {
	modifiedStr := strings.TrimSpace(*str)
	str = &modifiedStr
}
