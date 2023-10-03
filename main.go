package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateSecretKey() (string, error) {
	key := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

func main() {
	secretKey, err := generateSecretKey()
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return
	}
	fmt.Println("Generated SECRET_KEY:", secretKey)
}
