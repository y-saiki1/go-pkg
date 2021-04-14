package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type BlockCipher interface {
	Encrypt(plain []byte) ([]byte, error)
	Decrypt(encrypted []byte) ([]byte, error)
}

// 規格: AES/CBC/PKCS7
type AesCbcPkcs7Cipher struct {
	block cipher.Block
}

func NewAesCbcPkcs7Cipher(key []byte) (*AesCbcPkcs7Cipher, error) {
	keyLen := len(key)
	if (keyLen != 16) && (keyLen != 24) && (keyLen != 32) {
		return nil, fmt.Errorf("illegal key length [%d]. key length for AES must be 128, 192, 256 bit", keyLen)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher block")
	}
	return &AesCbcPkcs7Cipher{
		block: block,
	}, nil
}

func (c *AesCbcPkcs7Cipher) pad(b []byte) []byte {
	padSize := aes.BlockSize - (len(b) % aes.BlockSize)
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(b, pad...)
}

func (c *AesCbcPkcs7Cipher) unpad(b []byte) []byte {
	padSize := int(b[len(b)-1])
	return b[:len(b)-padSize]
}

func (c *AesCbcPkcs7Cipher) Encrypt(plain []byte) ([]byte, error) {
	padded := c.pad(plain)
	encryptedText := make([]byte, aes.BlockSize+len(padded))
	iv := encryptedText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCEncrypter(c.block, iv)
	cbc.CryptBlocks(encryptedText[aes.BlockSize:], padded)
	encryptedTextBase64 := base64.StdEncoding.EncodeToString(encryptedText)
	return []byte(encryptedTextBase64), nil
}

func (c *AesCbcPkcs7Cipher) Decrypt(encryptedTextBase64 []byte) (string, error) {
	encryptedText, _ := base64.StdEncoding.DecodeString(string(encryptedTextBase64))
	if len(encryptedText) < aes.BlockSize {
		return "", fmt.Errorf("encryped text must be lognder than block size")
	} else if (len(encryptedText) % aes.BlockSize) != 0 {
		return "", fmt.Errorf("encryped text must be multiple of blocksize(128bit)")
	}

	iv := encryptedText[:aes.BlockSize]
	cbc := cipher.NewCBCDecrypter(c.block, iv)
	encryptedText = encryptedText[aes.BlockSize:]
	plain := make([]byte, len(encryptedText))

	cbc.CryptBlocks(plain, encryptedText)

	return string(c.unpad(plain)), nil
}
