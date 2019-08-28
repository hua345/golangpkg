package encrypt

import "testing"

func TestEncodeSHA256(t *testing.T) {
	name := "fang"
	t.Log(EncodeSHA256(name))
	result := "a57d374337978ca7ff450e4d6411113b7a349d47a6608afd72fcd7b8f3aab965"
	if EncodeSHA256(name) != result {
		t.Error("SHA256 Failed")
	}
}
