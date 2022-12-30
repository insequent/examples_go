package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers
	nullTimer := &time.Timer{}
	secondTimer := time.NewTimer(time.Second)

	// Ticker
	millisecondTicker := time.NewTicker(time.Millisecond)

	deci := 0
	for {
		select {
		case <-nullTimer.C: // This will never happen
			fmt.Println("Null timer ticked")
		case <-secondTimer.C:
			fmt.Println("1 second up! All done")
			return
		case t := <-millisecondTicker.C:
			if t.UnixMilli()%100 == 0 {
				// This can sometimes go over 10. Not entirely sure why...
				deci++
				fmt.Printf("%d decisecond has passed\n", deci)
			}
		}
	}
}
