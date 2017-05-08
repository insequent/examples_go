package main

import (
    "fmt"
)

type Blah struct {
    Things map[string]interface{}
}

func main() {
    b := Blah{}

    if b.Things == nil {
        fmt.Println("True")
    }
}
