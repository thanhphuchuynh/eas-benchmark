package main

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"
)

func GenNonce(size int) ([]byte, error) {
	b := make([]byte, size)

	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

var PlainText, _ = GenNonce(4096)

var AuthKey256, _ = GenNonce(32)
var EncKey256, _ = GenNonce(32)
var EncKey128, _ = GenNonce(16)
var CipherText, _ = Encrypt(EncKey128, string(PlainText))

const N = 1000000

func BenchmarkOneEncrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encrypt(EncKey128, string(PlainText))
	}
}
func BenchmarkOneDecrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Decrypt(EncKey128, CipherText)
	}
}
func BenchmarkOneEncryptDecrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		str, _ := Encrypt(EncKey128, string(PlainText))
		Decrypt(EncKey128, str)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	NaiveTimesEncrypt(N)
	for n := 0; n < b.N; n++ {
		elasped := NaiveTimesEncrypt(N)
		fmt.Printf("Run %d. Did %d iterations in %s\n", n, N, elasped)
	}
}
func BenchmarkDecrypt(b *testing.B) {
	NaiveTimesDecrypt(N)
	for n := 0; n < b.N; n++ {
		elasped := NaiveTimesDecrypt(N)
		fmt.Printf("Run %d. Did %d iterations in %s\n", n, N, elasped)
	}
}
func NaiveTimesEncrypt(n int) time.Duration {
	startTime := time.Now()

	for i := 0; i < n; i++ {
		encText, err := Encrypt(EncKey128, string(PlainText))

		if err != nil {
			panic(err)
		}

		if encText == "" {
			panic("Encrypted text cannot be nil")
		}
	}

	elapsed := time.Since(startTime)

	return elapsed
}

func NaiveTimesDecrypt(n int) time.Duration {
	startTime := time.Now()

	for i := 0; i < n; i++ {
		text, err := Decrypt(EncKey128, CipherText)

		if err != nil {
			panic(err)
		}

		if text == "" {
			panic("Plain text cannot be nil")
		}
	}

	elapsed := time.Since(startTime)

	return elapsed
}
