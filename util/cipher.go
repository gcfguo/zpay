package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strconv"
)

const _source = `0123456789`

type Signature struct {
	zPayMchID uint64
	appID     string
	appSecret string
}

func NewSignature(
	zPayMchID uint64,
	appID string,
	appSecret string) *Signature {
	return &Signature{
		zPayMchID: zPayMchID,
		appID:     appID,
		appSecret: appSecret,
	}
}

func (s *Signature) Sign() string {
	core := sha256.New()
	b1 := []byte(s.appID)
	b2 := []byte(s.appSecret)
	b3 := s.Salt()
	_, _ = core.Write(b1)
	_, _ = core.Write(b2)
	_, _ = core.Write(b3)
	return fmt.Sprintf("%x", md5.Sum(core.Sum(nil)))
}

func (s *Signature) Salt() []byte {
	sz := strconv.FormatUint(s.zPayMchID, 10)
	var buf bytes.Buffer
	for _, v := range sz {
		b := []byte{_source[v-'0']}
		_, _ = buf.Write(b)
	}
	return buf.Bytes()
}

type Cipher struct {
	zPayMchID uint64
	appID     string
	appSecret string
}

func NewCipher(
	zPayMchID uint64,
	appID string,
	appSecret string,
) *Cipher {
	return &Cipher{
		zPayMchID: zPayMchID,
		appID:     appID,
		appSecret: appSecret,
	}
}

func (c *Cipher) Encrypt(plaintext string) (string, error) {
	sg := NewSignature(c.zPayMchID, c.appID, c.appSecret)
	secretKey := sg.Sign()
	ec := NewAesPKCS7(secretKey)
	ciphertext, err := ec.Decrypt(plaintext)
	if err != nil {
		return "", err
	}
	return ciphertext, nil
}

func (c *Cipher) Decrypt(ciphertext string) (string, error) {
	sg := NewSignature(c.zPayMchID, c.appID, c.appSecret)
	secretKey := sg.Sign()
	ec := NewAesPKCS7(secretKey)
	plaintext, err := ec.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}
	return plaintext, nil
}
