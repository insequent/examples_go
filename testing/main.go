package main

import (
	"fmt"
	"math/rand"
)

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += string(LETTERS[rand.Intn(len(LETTERS))])
	}
	return result
}

func randomBytes(length int) []byte {
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, LETTERS[rand.Intn(len(LETTERS))])
	}
	return result
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(randomString(i))
	}
}
