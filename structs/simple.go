package main

import (
	"fmt"
)

type Blah struct {
	Maps  map[string]interface{}
	Array []string
	Int   int
}

func NewBlah() *Blah {
	return &Blah{
		Maps:  map[string]interface{}{},
		Array: []string{},
	}
}

func main() {
	a := NewBlah()
	b := NewBlah()

	a.Maps["Test1"] = "This is the first test"
	b.Maps = a.Maps
	a.Maps["Test2"] = "This is the second test"

	b.Array = []string{"a", "b", "c"}
	a.Array = b.Array

	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Printf("Int: %d\n", a.Int)
}
