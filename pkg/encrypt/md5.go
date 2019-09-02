package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(value))

	return hex.EncodeToString(md5Hash.Sum(nil))
}
