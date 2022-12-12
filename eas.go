package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(key []byte, message string) (string, error) {

	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed new cipher %v", err)
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed gen iv %v", err)
	}

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	//returns to base64 encoded string
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key []byte, securemess string) (string, error) {

	cipherText, err := base64.StdEncoding.DecodeString(securemess)
	if err != nil {
		return "", fmt.Errorf("failed DecodeString %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed new cipher %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipherText block size is too short")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
