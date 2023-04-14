package main

import (
	"fmt"
)

func addToSlice(slice *[]string, str string) {
	*slice = append(*slice, str)
}

func main() {
	blah := []string{
		"this",
		"that",
		"them",
		"they",
	}
	fmt.Println("Slice:", blah)

	fmt.Println("Adding 'more!' to the slice...")
	blah = append(blah, "more!")
	fmt.Println("Slice: ", blah)

	// Limited helpers
	fmt.Println("Removing 'that' from slice...")
	index := -1
	for i, e := range blah {
		if e == "that" {
			index = i
		}
	}
	if index >= 0 {
		blah = append(blah[:index], blah[index+1:]...)
	}

	fmt.Println("Slice:", blah)
	for i, s := range blah {
		fmt.Printf("index %d: %v\n", i, s)
	}

	// Copying a slice must use copy() to avoid copying by reference (shallow copy)
	fmt.Println("Making new slice test2...")
	test2 := make([]string, len(blah))
	copy(test2, blah)
	test2 = test2[:1]
	test2 = append(test2, "For test 2")

	fmt.Println("Test2:", test2)
	fmt.Println("Blah:", blah)

	// Slice elements are passed by reference, but the slice as a whole is not.
	// You need to pass a pointer if you want to change a slice by function call
	test3 := []string{"a", "b", "c"}
	addToSlice(&test3, "d")
	fmt.Println("Test3:", test3)
}
