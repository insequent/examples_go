package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBlue(t *testing.T) {
	fmt.Println("Running TestBlue")
	os.Setenv("TEST_VAR", "blue")
	printVar()
}

func TestDefault(t *testing.T) {
	fmt.Println("Running TestDefault")
	printVar()
}
