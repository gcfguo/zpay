package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesPKCS7 struct {
	key string
}

func NewAesPKCS7(key string) *AesPKCS7 {
	return &AesPKCS7{key: key}
}

func (a *AesPKCS7) EncryptKey() []byte {
	return []byte(a.key)
}

func (a *AesPKCS7) DecryptKey() []byte {
	return []byte(a.key)
}

func (a *AesPKCS7) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(a.EncryptKey())
	if err != nil {
		return "", err
	}

	// 使用随机生成的 IV（初始向量）
	iv := make([]byte, aes.BlockSize)

	stream := cipher.NewCTR(block, iv)

	// 加密数据
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// PKCS#7 填充
	paddedCiphertext := a.pKCS7Padding(ciphertext, aes.BlockSize)

	// 返回进行Base64编码后的密文
	return base64.StdEncoding.EncodeToString(paddedCiphertext), nil
}

func (a *AesPKCS7) Decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(a.DecryptKey())
	if err != nil {
		return "", err
	}

	// 对 Base64 编码的密文进行解码
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// PKCS#7 去除填充
	ciphertext = string(a.pKCS7UnPadding(decodedCiphertext))

	iv := make([]byte, aes.BlockSize) // CTR 模式中 IV 为全零
	stream := cipher.NewCTR(block, iv)

	// 解密数据
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, []byte(ciphertext))

	return string(plaintext), nil
}

// pKCS7Padding 使用 PKCS#7 填充数据
func (a *AesPKCS7) pKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padData...)
}

// pKCS7UnPadding 去除 PKCS#7 填充
func (a *AesPKCS7) pKCS7UnPadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}
