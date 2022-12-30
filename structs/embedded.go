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

func (p *Person) whisper() {
	fmt.Println("Are children even people?")
}

type Child struct {
	Name string
	*Person
}

func (c *Child) SayHi() {
	fmt.Println("Hi! Do you like dinosaurs? My mom hates eggs. You smell like purple")
}

func (c *Child) SayName() {
	fmt.Printf("My name is %s, and I hate authority\n", c.Name)
}

func (c *Child) whisper() {
	fmt.Println("Hey?! What are whispering about??")
}

func main() {
	child := &Child{"Chet", &Person{}}
	child.SayName()
	child.SayHi()
	child.whisper()
}
