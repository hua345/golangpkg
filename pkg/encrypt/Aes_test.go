package encrypt

import (
	"encoding/base64"
	"github.com/hua345/golangpkg/pkg/util"
	"os/exec"
	"strings"
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

func TestAESDecryptCBC(t *testing.T) {
	uuid := util.GetUUID32()
	t.Log("uuid:", uuid)
	var aesKey = []byte(uuid)
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
	c := exec.Command("cmd", "/C", "node", "cryptojsTest.js", string(data), string(aesKey))
	result, err := c.Output()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(strings.TrimSpace(string(result)))
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
