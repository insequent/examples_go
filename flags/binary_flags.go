package main

import (
	"fmt"
)

const (
	first = 1 << iota
	second
	third
	fourth
	fifth
	sixth
	seventh
	eighth
	ninth
)

func main() {
	fmt.Printf("first (%T): %b\n", first, first)
	fmt.Printf("second (%T): %b\n", second, second)
	fmt.Printf("third (%T): %b\n", third, third)
	fmt.Println("...")
	fmt.Printf("ninth (%T): %b\n", ninth, ninth)

	var flags uint8
	fmt.Printf("Initial flags (%T): %b\n", flags, flags)

	// Add first and second flags
	flags = first + second
	fmt.Printf("Adding first and second flag: %b\n", flags)

	// Add third and subtract second flags
	flags = flags + third - second
	fmt.Printf("Adding third and subtracting second: %b\n", flags)

	// Check if second flag is unset
	if flags&second != second {
		fmt.Println("Second flag is not set")
	}

	// Check if third flag is set
	if flags&third == third {
		fmt.Println("Third flag is set")
	}

	// Simultaneously add fourth and fifth flag
	newFlags := uint8(fourth + fifth)
	flags += newFlags

	// Check if fourth and fifth flag are set
	if flags&newFlags == newFlags {
		fmt.Println("Fourth and Fifth flag are set")
	}

	// Print final result
	fmt.Printf("Final flags: %b\n", flags)

	// Add ninth flag (OVERFLOW, just playing here...)
	//flags = flags + ninth
	//fmt.Printf("Adding ninth flag: %b\n", flags)
}
