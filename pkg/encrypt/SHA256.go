package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncodeSHA256(value string) string {
	m := sha256.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
