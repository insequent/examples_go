package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's loop!")
	fmt.Println()

	// 2D array iteration
	nums := [][]uint8{
		{0, 1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, 6, 7, 0},
		{2, 3, 4, 5, 6, 7, 0, 1},
		{3, 4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2, 3},
		{5, 6, 7, 0, 1, 2, 3, 4},
		{6, 7, 0, 1, 2, 3, 4, 5},
		{7, 0, 1, 2, 3, 4, 5, 6},
	}

	// regular sort
	fmt.Println("Counting up!")
	for i := 0; i <= 7; i++ {
		for _, num := range nums[i] {
			fmt.Printf(" %v", num)
		}
		fmt.Println()
	}

	fmt.Println("Counting down!")
	for i := 7; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			fmt.Printf(" %v", nums[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	// Nil map fun
	var i interface{}
	i = nil

	fmt.Println("Let's see what happens when we loop a nil map[string]interface")
	if m, ok := i.(map[string]interface{}); ok {
		for k, v := range m {
			fmt.Println("Shocking! We're inside a nil loop!", k, v)
		}
	} else {
		fmt.Println("Eek, we're nil! No loop for us")
	}

	fmt.Printf("The type of an nil map is %T\n", i)
	fmt.Println()

	// For + Select break
	intCh := make(chan int)
	go func() {
		for i := 0; i <= 5; i++ {
			intCh <- i
		}
		fmt.Println("Worker done!")
	}()

loop:
	for {
		select {
		case i := <-intCh:
			if i == 5 {
				break loop
			}
			fmt.Println("Select loop received:", i)
		}
	}
	fmt.Println("We've broken out of the select loop!")
}
