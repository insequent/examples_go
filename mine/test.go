package main

import (
	"fmt"
)

func IsEven(x int) bool {
	return x%2 == 0
}

func main() {
	num := 3

	if !IsEven(num) {
		fmt.Println("NOT EVEN!")
	}
}
