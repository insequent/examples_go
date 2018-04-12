package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "/home/dark/go"

	fmt.Println("Dir:", filepath.Dir(path))
	fmt.Println("Base:", filepath.Base(path))
	fmt.Println("Join:", filepath.Join(path, "test.txt"))
}
