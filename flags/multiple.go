package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	credentialFlags := flag.NewFlagSet("Credentials", flag.ContinueOnError)

	username := credentialFlags.String("username", "user", "Your username")
	password := credentialFlags.String("password", "password", "Your password")

	commandFlags := flag.NewFlagSet("Commands", flag.ContinueOnError)

	create := commandFlags.Bool("create", false, "Creates something")
	read := commandFlags.Bool("read", false, "Reads something")
	update := commandFlags.Bool("update", false, "Updates something")
	del := commandFlags.Bool("delete", false, "Deletes something")

	resourceFlags := flag.NewFlagSet("Resources", flag.ContinueOnError)

	file := resourceFlags.String("file", "", "File name")
	pipe := resourceFlags.String("pipe", "", "Pipe name")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [credential options...] [command option] [resource option]\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Perform a command against a resource! Fun!\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nCredential Options:\n")
		credentialFlags.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nCommand Options:\n")
		commandFlags.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nResource Options:\n")
		resourceFlags.PrintDefaults()
	}

	flag.Parse()
	credentialFlags.Parse(os.Args[1:])
	commandFlags.Parse(os.Args[1:])
	resourceFlags.Parse(os.Args[1:])

	fmt.Println("Args:", os.Args)
	fmt.Println()
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println()
	fmt.Println("create:", *create)
	fmt.Println("read:", *read)
	fmt.Println("update:", *update)
	fmt.Println("delete:", *del)
	fmt.Println()
	fmt.Println("file:", *file)
	fmt.Println("pipe:", *pipe)
}
