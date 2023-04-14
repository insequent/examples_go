package main

import (
	"encoding/json"
	"log"
)

type SuperSet struct {
	A string
	B string
	C string
}

type SubSet struct {
	A string
}

type UnrelatedSet struct {
	Z string
}

func main() {
	data := []byte(`{"A":"Something A", "B":"Something B", "C":"Something C"}`)

	super := &SuperSet{}
	if err := json.Unmarshal(data, super); err != nil {
		log.Fatal("Error unmarshalling into SuperSet:", err)
	}

	log.Println(super)

	sub := &SubSet{}
	if err := json.Unmarshal(data, sub); err != nil {
		log.Fatal("Error unmarshalling into SubSet:", err)
	}

	// Note that unused fields are ignored
	log.Println(sub)

	unrelated := &UnrelatedSet{}
	if err := json.Unmarshal(data, unrelated); err != nil {
		log.Fatal("Error unmarshalling into UnrelatedSet:", err)
	}

	// No fields match here, so we end up with an empty struct
	log.Println(unrelated)
}
