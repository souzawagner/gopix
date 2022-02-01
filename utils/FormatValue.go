package utils

import (
	"fmt"
	"regexp"
)

func FormatValue(id, value string) string {
	value = CleanString(value)
	size := fmt.Sprintf("%02d", len(value)) // pad zeroes
	return fmt.Sprintf("%s%s%s", id, size, value)
}

func CleanString(value string) string {
	return regexp.
		MustCompile(`[^A-z0-9$%*+-./:\s]`).
		ReplaceAllString(value, ``)
}
