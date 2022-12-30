package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
	Type() string
}

type Dog struct {
}

func (d *Dog) Speak() string {
	return "Woof!"
}

func (d *Dog) Type() string {
	return "Dog"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func (c *Cat) Type() string {
	return "Cat"
}

type Llama struct {
}

func (l *Llama) Speak() string {
	return "?????"
}

func (l *Llama) Type() string {
	return "Llama"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func (j *JavaProgrammer) Type() string {
	return "JavaProgrammer"
}

type Horse struct {
}

func (h Horse) Speak() string {
	return "Nay"
}

func (h Horse) Type() string {
	return "Horse"
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}, &Llama{}, &JavaProgrammer{}, &Horse{}}
	for _, animal := range animals {
		fmt.Printf("%s goes %q\n", animal.Type(), animal.Speak())
	}
}
