package main

import (
	"bytes"
	"fmt"
)

func main() {
	w := bytes.Buffer{}
	w.Write([]byte("Hi!"))

	fmt.Println(string(w.Bytes()))
}
