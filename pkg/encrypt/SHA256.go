package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncodeSHA256(value string) string {
	shaHash := sha256.New()
	shaHash.Write([]byte(value))

	return hex.EncodeToString(shaHash.Sum(nil))
}
