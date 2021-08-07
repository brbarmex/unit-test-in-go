package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"os"
)

func Encrypt(value string) string {
	key := []byte(os.Getenv("secret_aes"))
	plaintext := []byte(value)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte(os.Getenv("secret_aes_cms"))
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(value string) string {
	key := []byte(os.Getenv("secret_aes"))
	ciphertext, _ := hex.DecodeString(value)
	nonce := []byte(os.Getenv("secret_aes_cms"))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}
