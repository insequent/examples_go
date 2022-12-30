package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	talkCh := make(chan struct{})

	Say := func(talkCh chan struct{}, i int, wg *sync.WaitGroup) {
		defer wg.Done()

		timer := time.NewTimer(time.Minute * time.Duration(1))

		select {
		case <-talkCh:
			fmt.Printf("%v:\tI'm saying something!\n", i)
		case <-timer.C:
			return
		}
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Say(talkCh, i, &wg)
	}

	for i := 1; i <= 10; i++ {
		talkCh <- struct{}{}
	}

	wg.Wait()
}
