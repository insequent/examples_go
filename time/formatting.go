package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("%11s: %d\n", "Unix", now.Unix())
	fmt.Printf("%11s: %s\n", "RFC822", now.Format(time.RFC822))
	fmt.Printf("%11s: %s\n", "RFC822Z", now.Format(time.RFC822Z))
	fmt.Printf("%11s: %s\n", "RFC850", now.Format(time.RFC850))
	fmt.Printf("%11s: %s\n", "RFC1123", now.Format(time.RFC1123))
	fmt.Printf("%11s: %s\n", "RFC1123Z", now.Format(time.RFC1123Z))
	fmt.Printf("%11s: %s\n", "RFC3339", now.Format(time.RFC3339))
	fmt.Printf("%11s: %s\n", "RFC3339Nano", now.Format(time.RFC3339Nano))

	fmt.Printf("%11s: %s\n", "Custom", now.Format("2006-01-02 15:04Z07"))
	fmt.Printf("%11s: %s\n", "Custom UTC", now.UTC().Format("2006-01-02 15:04-07"))
}
