package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Child1 struct {
	Name string
}

type Child2 struct {
	Name string
}

type Parent struct {
	HasChildren bool
	Child1
	Child2
}

func main() {
	parent := Parent{
		HasChildren: true,
		Child1: Child1{
			Name: "child1",
		},
		Child2: Child2{
			Name: "child2",
		},
	}

	b, err := json.MarshalIndent(parent, "", "\t")
	if err != nil {
		log.Fatalf("Error while marshalling parent: %v", err)
	}

	// WARN: Notice the child values are lost since their field names match
	fmt.Println(string(b))
}
