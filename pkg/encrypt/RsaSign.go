package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func RsaSignWithSha1Base64(originalData []byte, privateKey []byte) (string, error) {
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
	h := sha1.New()
	h.Write(originalData)
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	signData := base64.StdEncoding.EncodeToString(signature)
	return signData, nil
}

func RsaVerySignWithSha1Base64(signData string, originalData, publicKey []byte) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hash.Sum(nil), sign)
}
