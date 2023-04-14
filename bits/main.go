package main

import "fmt"

// Bits are a convenient way to store flags/options
const (
	A_FLAG int = 1 << iota
	B_FLAG
	C_FLAG
	D_FLAG
	E_FLAG
	F_FLAG
)

func main() {
	tests := []int{
		A_FLAG | B_FLAG | C_FLAG,
		A_FLAG | F_FLAG,
		D_FLAG,
	}

	for i, test := range tests {
		fmt.Printf("Test %d:\n", i)
		if test&A_FLAG != 0 {
			fmt.Println("\tFlag A is set")
		}
		if test&B_FLAG != 0 {
			fmt.Println("\tFlag B is set")
		}
		if test&C_FLAG != 0 {
			fmt.Println("\tFlag C is set")
		}
		if test&D_FLAG != 0 {
			fmt.Println("\tFlag D is set")
		}
		if test&E_FLAG != 0 {
			fmt.Println("\tFlag E is set")
		}
		if test&F_FLAG != 0 {
			fmt.Println("\tFlag F is set")
		}
	}
}
