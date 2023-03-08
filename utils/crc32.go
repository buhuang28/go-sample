package utils

import "hash/crc32"

func CRC32(b []byte) uint32 {
	return crc32.ChecksumIEEE(b)
}
