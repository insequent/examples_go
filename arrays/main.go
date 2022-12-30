package main

import (
	"fmt"
)

func main() {
	blah := make([]string, 5)
	blah[0] = "this"
	blah[1] = "that"
	blah[2] = "them"
	blah[3] = "they"
	blah[4] = "more!"

	fmt.Println("Array:", blah)

	for i, element := range blah {
		fmt.Printf("Array index %d: %s\n", i, element)
	}

	// Not a whole lot to do with arrays...
}
