package encrypt

import "testing"

func TestEncodeHmacSha256(t *testing.T) {
	name := "fang"
	key := "love"
	t.Log(EncodeHmacSha256(name, key))
	result := "48ff7d2604e73afd6f439d9aa79bfb4a937980e394fe6147327b0824212b3c48"
	if EncodeHmacSha256(name, key) != result {
		t.Error("EncodeHmacSha256 Failed")
	}
}
