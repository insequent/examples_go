package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	config := *flag.String("c", "", "Absolute path to the config file")
	flag.Parse()

	if config == "" {
		fmt.Println("ERROR: Script requires a config file path")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Config path is", config)

	args := flag.Args()
	fmt.Printf("Received the following args without flags: %v\n", args)
}
