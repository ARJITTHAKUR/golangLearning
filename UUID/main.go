package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length  = 20 // Change the length to your desired length
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomString := generateRandomString(length)
	fmt.Println(randomString)
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
