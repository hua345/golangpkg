package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 加密
func RsaEncrypt(orignData []byte, publicKey []byte) (string, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	encryptData, err := rsa.EncryptPKCS1v15(rand.Reader, pub, orignData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptData), nil
}

// 解密
func RsaDecrypt(base64EncryptedData string, privateKey []byte) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(base64EncryptedData)
	if err != nil {
		return "", errors.New("Decode base64 error!")
	}
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	decryptData, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encryptedData)
	if err != nil {
		return "", err
	}
	return string(decryptData), nil
}
