package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func EncodeHmacSha256(value, key string) string {
	macHash := hmac.New(sha256.New, []byte(key))
	macHash.Write([]byte(value))
	return hex.EncodeToString(macHash.Sum(nil))
}
