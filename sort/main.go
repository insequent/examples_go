package main

import (
	"fmt"
	"sort"
)

type Color struct {
	Name string
}

// Colors is a slice of Color implementing Len(), Swap(a, b), and Less(a, b) for sorting
type Colors []Color

func (cs Colors) Len() int {
	return len(cs)
}

func (cs Colors) Swap(A, B int) {
	cs[A], cs[B] = cs[B], cs[A]
}

func (cs Colors) Less(A, B int) bool {
	return cs[A].Name < cs[B].Name
}

func main() {
	// Strings
	strings := []string{"a", "c", "f", "b", "e", "h", "g"}
	sort.Strings(strings)
	fmt.Println(strings)

	// Ints
	ints := []int{3, 1, 2, 7, 5, 4, 6}
	sort.Ints(ints)
	fmt.Println(ints)

	// Structs with methods
	colors := Colors{{Name: "red"}, {Name: "blue"}, {Name: "yellow"}}
	sort.Sort(colors)
	fmt.Println(colors)

	// Structs without methods
	animals := []struct{ Name string }{{Name: "hamster"}, {Name: "squirrel"}, {Name: "mouse"}}
	sort.SliceStable(animals, func(A, B int) bool { return animals[A].Name < animals[B].Name })
	fmt.Println(animals)
}
