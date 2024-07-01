package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run password-generator.go <length>")
		return
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil || length <= 0 {
		fmt.Println("Invalid length specified. Length must be a positive integer.")
		return
	}

	fmt.Println(randomString(length))
}

func randomString(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // handle error properly
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:length]
}
