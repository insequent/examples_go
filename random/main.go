package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(1234)
	round1 := make([]int, 5)
	for i := 0; i < 5; i += 1 {
		round1[i] = rand.Int()
	}
	fmt.Println("First round with seed 123:", round1)

	rand.Seed(1234)
	round2 := make([]int, 5)
	for i := 0; i < 5; i += 1 {
		round2[i] = rand.Int()
	}
	fmt.Println("Second round with seed 123:", round2)

	equal := true
	for i := 0; i < 5; i += 1 {
		if round1[i] != round2[i] {
			equal = false
			break
		}
	}

	fmt.Println("Are they equal?", equal)

	rand.Seed(time.Now().Unix())
	round3 := make([]int, 5)
	for i := 0; i < 5; i += 1 {
		round3[i] = rand.Int()
	}
	fmt.Println("Third round using time.Now() as seed:", round3)

	equal = true
	for i := 0; i < 5; i += 1 {
		if round1[i] != round3[i] {
			equal = false
			break
		}
	}

	fmt.Println("Is third round the same as round one?", equal)

}
