package utils

import "strconv"

func Dechex(number int64) string {
	return strconv.FormatInt(number, 16)
}
