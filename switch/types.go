package main

import "fmt"

type Dog struct{}
type Cat struct{}
type Bird struct{}

func printType(i interface{}) {
	switch i.(type) {
	case Dog:
		fmt.Println("Woof")
	case Cat:
		fmt.Println("Meow")
	case Bird:
		fmt.Println("Tweet Tweet")
	}
}

func main() {
	printType(Bird{})
	printType(Cat{})
	printType(Dog{})
}
