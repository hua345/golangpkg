package encrypt

import "testing"

func TestPasswordEncrypt(t *testing.T) {
	name := "fang"
	t.Log(PasswordEncrypt(name))
	result := "a111b32bee69055cc0021660dc11ace9"
	if PasswordEncrypt(name) != result {
		t.Error("PasswordEncrypt Failed")
	}
}
