package encrypt

import (
	"testing"
)

// Rsa加密解密测试
func TestRsaEncrypt(t *testing.T) {
	//原始数据
	var data = []byte("fang")
	//rsa加密
	encryptData, err := RsaEncrypt(data, PublicKey)
	if err != nil {
		t.Log("encrypt err: ", err)
		return
	}
	t.Log("encryptData:", encryptData)
	//rsa解密
	decryptData, err := RsaDecrypt(encryptData, PrivateKey)
	if err != nil {
		t.Log("decrypt err:", err)
		return
	}
	t.Log("decryptData:", decryptData)

	if decryptData != string(data) {
		t.Error("RSA 加解密失败")
	}
}

// RSA加签验签测试
func TestRsaSign(t *testing.T) {
	//原始数据
	var data = []byte("fang")
	//加签
	signData, err := RsaSignWithSha1Base64(data, PrivateKey)
	if err != nil {
		t.Log("sign err:", err)
		return
	}
	t.Log("signData:", signData)
	//验签
	verifyResult := RsaVerySignWithSha1Base64(signData, data, PublicKey)
	if verifyResult != nil {
		t.Log("verify failed:", verifyResult)
		return
	}
	t.Log("verifyResult:", verifyResult)

	if verifyResult != nil {
		t.Error("RSA 验签失败")
	}
}
