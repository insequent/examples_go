package main

import "fmt"

func main() {
	str := "This is a string"
	i := 4
	blah := map[string]string{
		"this":  "test",
		"is":    "test",
		"great": "test?",
	}

	fmt.Println("Println:", str, i)
	fmt.Printf("Printf: %v %v\n", str, i)
	fmt.Printf("String: %s\n", blah)
	fmt.Printf("Default: %v\n", blah)
	fmt.Printf("Type: %T\n", blah)

	longStr := "This is a randomly long string..."
	fmt.Printf("%.8s\n", longStr)
}
