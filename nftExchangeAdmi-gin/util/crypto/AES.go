package crypto

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

// PKCS5Padding 实现 PKCS5 填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS5UnPadding 实现 PKCS5 去填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// AesEncryptECB AES ECB 加密（传入和传出字符串）
func AesEncryptECB(plaintext string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// PKCS5Padding 填充
	paddedText := PKCS5Padding([]byte(plaintext), block.BlockSize())

	ciphertext := make([]byte, len(paddedText))

	// ECB 模式加密每个块
	for start := 0; start < len(paddedText); start += block.BlockSize() {
		block.Encrypt(ciphertext[start:start+block.BlockSize()], paddedText[start:start+block.BlockSize()])
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesDecryptECB AES ECB 解密（传入和传出字符串）
func AesDecryptECB(ciphertext string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(decodedCiphertext))

	// ECB 模式解密每个块
	for start := 0; start < len(decodedCiphertext); start += block.BlockSize() {
		block.Decrypt(plaintext[start:start+block.BlockSize()], decodedCiphertext[start:start+block.BlockSize()])
	}

	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext), nil
}
