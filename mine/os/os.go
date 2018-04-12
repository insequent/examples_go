package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "os.go"

	stats, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stats.Name())
	fmt.Println(stats.Size())
	fmt.Println(stats.Mode())
	fmt.Println(stats.ModTime())
	fmt.Println(stats.IsDir())
	fmt.Println(stats.Sys())

}
