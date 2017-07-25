package main

import (
	"fmt"
)

func ret_num(x int) int {
	return x
}

func add_one(x int) int {
	y := x + 1
	return y
}

func main() {
	fmt.Println(add_one(ret_num(4)))
}
