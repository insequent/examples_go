package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	err := enc.Encode(struct {
		Foo string `json:"foo"`
	}{
		Foo: "something",
	})

	fmt.Println("JSON:", buf.String())
	fmt.Println("err:", err)
}
