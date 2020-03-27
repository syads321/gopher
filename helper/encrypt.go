package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

const key = "6368616e676520746869732070617373776f726420746f206120736563726574"
const nonce = "64a9433eae7ccceee2fc0eda"

// Encrypt this
func Encrypt(text string) (string, error) {
	hexkey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	hexnonce, err := hex.DecodeString(nonce)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(hexkey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	bytetext := []byte(text)
	reciphertext := aesgcm.Seal(nil, hexnonce, bytetext, nil)
	if err != nil {
		panic(err.Error())
	}
	return hex.EncodeToString(reciphertext), err
}

// Decrypt fdkfld
func Decrypt(keystring string) (string, error) {

	hexkey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(keystring)
	if err != nil {
		return "", err
	}
	hexnonce, err := hex.DecodeString(nonce)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(hexkey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, hexnonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), err
}
