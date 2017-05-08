// hello.go

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    var who []string

    if len(os.Args) > 1 {
        who = strings.Split(os.Args[1], ",")
	} else {
        who = []string{"World!"}
    }

    for _, name := range who {
		fmt.Println("Hello", name)
	}

    m := map[string]string{
        "stuff": "things",
        "more":  "less",
        "know":  "not",
    }

    fmt.Printf("Map looks like %s", m)
}
