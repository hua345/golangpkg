package encrypt

import (
	"encoding/base64"
	"testing"
)

func TestAES(t *testing.T) {
	var aeskey = []byte("321423u9y8d2fwfl")
	data := []byte("fang")
	xpass, err := AesEncrypt(data, aeskey)
	if err != nil {
		t.Log(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	t.Log("AES加密后: ", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		t.Log(err)
		return
	}

	decryptData, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("AES解密后: ", string(decryptData))
	if string(data) != string(decryptData) {
		t.Error("AES加解密失败")
	}
}
