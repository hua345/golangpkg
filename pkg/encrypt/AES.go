package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesEncryptCBC(origData, key []byte) ([]byte, error) {
	// NewCipher creates and returns a new cipher.Block.The key argument should be the AES key
	// 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	iv := key[:blockSize]
	// 数据补齐/填充（Padding）
	origData = PKCS5Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cipherText, origData)
	return cipherText, nil
}

func AesDecryptCBC(cipherText, key []byte) ([]byte, error) {
	// NewCipher creates and returns a new cipher.Block.The key argument should be the AES key
	// 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	iv := key[:blockSize]
	// 解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	// 解密
	blockMode.CryptBlocks(origData, cipherText)
	// 去除填充
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
func AesEncryptCFB(origData, key []byte) ([]byte, error) {
	// NewCipher creates and returns a new cipher.Block.The key argument should be the AES key
	// 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	iv := key[:blockSize]
	// 数据补齐/填充（Padding）
	origData = PKCS5Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(origData))
	// 加密
	blockMode.XORKeyStream(cipherText, origData)
	return cipherText, nil
}

func AesDecryptCFB(cipherText, key []byte) ([]byte, error) {
	// NewCipher creates and returns a new cipher.Block.The key argument should be the AES key
	// 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	iv := key[:blockSize]
	// 解密模式
	blockMode := cipher.NewCFBDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	// 解密
	blockMode.XORKeyStream(origData, cipherText)
	// 去除填充
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
