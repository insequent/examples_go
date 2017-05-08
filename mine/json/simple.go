package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	A string
	B string
	C string
}

var in []byte = []byte(`{"A": "1", "C": "3"}`)

func main() {
	out := Test{
		A: "default a",
		B: "default b",
		// C will be ""
	}
	fmt.Printf("Original 'out': %+v\n", out)
	fmt.Printf("'in': %s\n", in)
	err := json.Unmarshal(in, &out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New 'out': %#v\n", out)
}
