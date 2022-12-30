package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	A string
	B string
	C string
	D string
}

func main() {
	in := []byte(`{"A": "1", "C": "3"}`)
	fmt.Printf("'in': %s\n", in)
	out := Test{
		A: "default a",
		B: "default b",
	}
	fmt.Printf("'out': %+v\n", out)
	err := json.Unmarshal(in, &out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshalled: %+v\n", out)

	m, _ := json.Marshal(out)
	fmt.Printf("Marshalled: %+v\n", string(m))
}
