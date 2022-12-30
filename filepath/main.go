package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()

	fmt.Println("Dir:", filepath.Dir(path))
	fmt.Println("Base:", filepath.Base(path))
	fmt.Println("Join:", filepath.Join(path, "main.go"))
}
