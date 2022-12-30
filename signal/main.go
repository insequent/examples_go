package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// Block until a signal is received.
	s := <-c
	switch s {
	default:
		fmt.Printf("Unexpected signal received %q (%#v)\n", s, s)
	case os.Interrupt:
		fmt.Printf("Received interrupt signal: %q (%#v)\n", s, s)
	case syscall.SIGHUP:
		fmt.Printf("Received HUP signal: %q (%#v)\n", s, s)
	}
}
