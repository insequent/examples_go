package main

import "fmt"

func main() {
	blah := []string{
		"this",
		"that",
		"paddywack",
		"give a dog a bone",
	}

	for s := range blah {
		fmt.Println(s)
	}
}
