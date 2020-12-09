package encrypt

import (
	"encoding/base64"
	"testing"
)

func TestAESCBC(t *testing.T) {
	var aesKey = []byte("fangfangfangfang")
	data := []byte("fangfang")
	cipherText, err := AesEncryptCBC(data, aesKey)
	if err != nil {
		t.Log(err)
		return
	}

	cipherTextBase64 := base64.StdEncoding.EncodeToString(cipherText)
	t.Log("AES加密后: ", cipherTextBase64)

	bytesPass, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		t.Log(err)
		return
	}

	decryptData, err := AesDecryptCBC(bytesPass, aesKey)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("AES解密后: ", string(decryptData))
	if string(data) != string(decryptData) {
		t.Error("AES加解密失败")
	}
}

func TestAESCFB(t *testing.T) {
	var aesKey = []byte("fangfangfangfang")
	data := []byte("fangfang")
	cipherText, err := AesEncryptCFB(data, aesKey)
	if err != nil {
		t.Log(err)
		return
	}

	cipherTextBase64 := base64.StdEncoding.EncodeToString(cipherText)
	t.Log("AES加密后: ", cipherTextBase64)

	bytesPass, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		t.Log(err)
		return
	}

	decryptData, err := AesDecryptCFB(bytesPass, aesKey)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("AES解密后: ", string(decryptData))
	if string(data) != string(decryptData) {
		t.Error("AES加解密失败")
	}
}
