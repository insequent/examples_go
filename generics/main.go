package main

import "fmt"

// 1. Limit types inline
func AddStrings[S string | []byte](strs ...S) S {
	result := ""
	for _, str := range strs {
		result += string(str)
	}

	return S(result)
}

// 2. Limit types by definition
type Number interface {
	int | uint | float64
}

func AddNumbers[N Number](nums ...N) N {
	var result N
	for _, num := range nums {
		result += num
	}

	return result
}

func main() {
	fmt.Println("Adding strings:", AddStrings("A", "B", "C"))
	fmt.Println("Adding bytes:", string(AddStrings([]byte("A"), []byte("B"), []byte("C"))))
	fmt.Println("Note that we cannot combine strings and bytes in same function call")
	fmt.Println()

	fmt.Println("Adding integers:", AddNumbers(1, 2, 3))
	fmt.Println("Adding uints:", AddNumbers(uint(1), uint(2), uint(3)))
	fmt.Println("Adding floats:", AddNumbers(1.1, 2.2, 3.3))
	fmt.Println("Again, note that we cannot combine types, due to type differences.")
	fmt.Println("However, inferred types can allow us to do this:", AddNumbers(1, uint(2), 3)) // Can be misleading
}
