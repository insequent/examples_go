package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "os.go"

	stats, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error?", err)
	}

	fmt.Println("Name:", stats.Name())
	fmt.Println("Size:", stats.Size())
	fmt.Println("Mode:", stats.Mode())
	fmt.Println("ModTime:", stats.ModTime())
	fmt.Println("IsDir:", stats.IsDir())
	fmt.Println("Sys (returns Any):", stats.Sys())
}
