package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 600
	blah := strconv.Itoa(num)
	fmt.Println(blah)

	str := "600"
	num, _ = strconv.Atoi(str)
	fmt.Println(num)
}
