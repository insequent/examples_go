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

type ByteTest struct {
	A []byte
}

type RawTest struct {
	A json.RawMessage
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

	fmt.Println("\nRaw test...\n")

	nested := []byte("{\"B\":\"text\"}")
	raw := &RawTest{
		A: nested,
	}

	fmt.Println("Marshaled:")
	b, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	b, err = json.MarshalIndent(raw.A, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("Single field marshaled:", string(b))
	fmt.Println("Straight string:", string(raw.A))

	fmt.Println("\nByte slice test...\n")
	bt := &ByteTest{
		A: nested,
	}

	fmt.Println("Marshaled:")
	b, err = json.MarshalIndent(bt, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	b, err = json.MarshalIndent(bt.A, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("Single field marshaled:", string(b))
	fmt.Println("Straight string:", string(bt.A))
}
