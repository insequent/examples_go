package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
}

func (cerr *CustomError) Error() string {
	return cerr.Message
}

func main() {
	err := errors.New("This is a test error")
	fmt.Println("Test error:", err)

	err = fmt.Errorf("This is another test error")
	fmt.Printf("%%s formatted error: %s\n%%v formatted error: %v\n", err, err)

	err = &CustomError{Message: "This is a custom error"}
	fmt.Printf("Customer error: %+v\n", err)
}
