package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t := time.Now()
	fmt.Println("Current time:", t)

	for n := 0; n <= 24; n++ {
		t = time.Now().UTC().Add(time.Duration(n) * time.Hour)
		fmt.Printf("Time +%d hours: %v\n", n, t)
		fmt.Printf("%s \n", t.Weekday())
		fmt.Println("(rounded):", t.Round(time.Duration(24)*time.Hour))
	}

	fmt.Printf("This took %v\n", time.Since(start))
}
