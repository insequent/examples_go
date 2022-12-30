package main

import (
	"fmt"
)

func main() {
	blah := "I'm a string!"
	fmt.Println(blah)
	strStuff(&blah)
	fmt.Println(blah)
}

func strStuff(str *string) {
	*str = "Something new"
}
