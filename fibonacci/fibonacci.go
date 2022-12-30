// Mostly stolen from https://tour.golang.org/concurrency/5

package main

import "fmt"

func fibWorker(c chan int, quit chan struct{}) {
	x := 0
	y := 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan struct{})

	go fibWorker(c, quit)

	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}

	quit <- struct{}{}
}
