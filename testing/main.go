package main

import (
	"fmt"
	"os"
)

func printVar() {
	v := os.Getenv("TEST_VAR")

	fmt.Printf("\tTEST_VAR value: %s\n", v)
}
