package utils

import (
	"github.com/howeyc/crc16"
	"strings"
)

func GetCRC16Checksum(str string) string {
	result := crc16.ChecksumCCITTFalse([]byte(str))
	return strings.ToUpper(Dechex(int64(result)))
}
