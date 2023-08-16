package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) SayName() {
	fmt.Printf("My name is %s\n", p.Name)
}

type Child struct {
	Thing string
	*Person
}

func (c *Child) SayHi() {
	fmt.Println("Hi! Do you like dinosaurs? My mom hates eggs. You smell like purple")
}

func (c *Child) SayName() {
	fmt.Printf("My name is %s, and I like %s\n", c.Name, c.Thing)
}

func main() {
	child := &Child{
		Thing:  "butterflies",
		Person: &Person{Name: "Chad"},
	}
	// After declaration, you can address fields of the embedded struct
	child.Name = "Chet"

	child.SayName()
	child.SayHi()
}
