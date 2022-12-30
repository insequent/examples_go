package main

import (
	"fmt"
	"regexp"
)

func main() {
	bytes := []byte("This is NOT a string")
	regex := "(?P<caps>[A-Z]+)"

	fmt.Printf("Matching for grouped caps in the following:\n\t%s\n", string(bytes))
	re := regexp.MustCompile(regex)
	matches := re.FindAllSubmatch(bytes, -1)
	for _, match := range matches {
		fmt.Printf("Match: %s\n", string(match[1]))
	}
}
