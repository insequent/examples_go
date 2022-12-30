package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

type A struct {
	This string
	That int
	More map[string]int
}

type B struct {
	Stuff int
	Blah  string
}

func Test(obj interface{}) (y string, err error) {
	bytes, err := yaml.Marshal(obj)
	y = string(bytes)
	return
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	a := &A{
		This: "this",
		That: 2,
		More: map[string]int{
			"to be":  1,
			"or not": 0,
		},
	}

	b := &B{
		Stuff: 37888,
		Blah:  "so much blah",
	}

	bytes, err := yaml.Marshal(a)
	CheckErr(err)
	fmt.Printf("'A' marshalled:\n%v\n", string(bytes))

	aYAML, err := Test(a)
	CheckErr(err)
	fmt.Printf("'A' interfaced and marshalled:\n%v\n", aYAML)

	bYAML, err := Test(b)
	CheckErr(err)
	fmt.Printf("'B' interfaced and marshalled:\n%v\n", bYAML)
}
