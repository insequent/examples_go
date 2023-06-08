package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Test struct {
	A string
	B string
	C string
}

func diffAny(any1, any2 any) string {
	return cmp.Diff(any1, any2)
}

func main() {
	test1 := Test{
		A: "A test",
		B: "B test",
		C: "C test",
	}

	test2 := Test{
		A: "A test",
		B: "something new",
		C: "C test",
	}

	fmt.Println(cmp.Diff(test1, test2))

	// Diff of interfaces
	fmt.Println(diffAny(test1, test2))

	// Diff of nil
	fmt.Println(diffAny(nil, test2))

	str1 := `
str := "This is fake code"
fmt.Println(str)
`

	str2 := `
str := "This is real code"
fmt.Println(str)
`

	fmt.Println(cmp.Diff(str1, str2))

	// Diff of nil and non-nillable
	fmt.Println(cmp.Diff(str1, nil))
}
