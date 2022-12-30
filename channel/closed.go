package main

import (
	"fmt"
	"time"
)

func main() {
	doneChan := make(chan int)

	ticker := time.NewTicker(time.Second)
	start := time.Now()
	end := start.Add(10 * time.Second)

	for {
		select {
		case <-doneChan:
			return
		case t := <-ticker.C:
			fmt.Println(t)
			if t.After(end) {
				close(doneChan)
			}
		}
	}
}
