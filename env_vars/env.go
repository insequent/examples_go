package main

import (
	"fmt"
	"os"
)

func main() {
	TEST := "test"
	fmt.Printf("TEST is at %p\n", &TEST)

	if test, ok := os.LookupEnv("TEST"); ok {
		fmt.Printf("test is at %p\n", &test)
		TEST = test
		fmt.Printf("Final TEST is at %p\n", &TEST)
	}

	// test var no longer exists
	fmt.Println("PWD:", os.Getenv("PWD"))
	fmt.Println("TEST:", TEST)
}
