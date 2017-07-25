package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Current time:", t)

	for n := 0; n <= 24; n++ {
		fmt.Printf("Time +%d hours: ", n)
		t = time.Now().UTC().Add(time.Hour * time.Duration(n))
		fmt.Printf("%s ", t.Weekday())
		t = time.Date(t.Round())
		fmt.Println(t)
	}
}
