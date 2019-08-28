package encrypt

import "testing"

func TestEncodeMD5(t *testing.T) {
	name := "fang"
	t.Log(EncodeMD5(name))
	result := "0b527ad35a7983fa5c9abdf31825c3cb"
	if EncodeMD5(name) != result {
		t.Error("Md5 Failed")
	}
}
