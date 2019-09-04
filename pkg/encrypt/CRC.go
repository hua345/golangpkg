package encrypt

import (
	"hash/crc32"
)

func CRC32Hash(value string) uint32 {
	crc32Hash := crc32.NewIEEE()
	_, err := crc32Hash.Write([]byte(value))
	if err != nil {
		panic(err)
	}
	return crc32Hash.Sum32()
}
