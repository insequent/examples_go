package main

import (
	"context"
	"fmt"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	cancel()
	// If we close the channel, then select with randomly select the closed channel and panic
	//close(ch)

	select {
	// Select won't send to ch without and active listener
	case ch <- struct{}{}:
		fmt.Println("Sent")
	case <-ctx.Done():
		fmt.Println("Context done")
	}
}
