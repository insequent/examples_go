package main

import (
	"fmt"
	"net/url"
)

func printURL(str string) {
	if u, err := url.Parse(str); err == nil {
		u.Query().Add("key", "value")
		fmt.Println("\tFragment:", u.Fragment)
		fmt.Println("\tScheme:", u.Scheme)
		fmt.Println("\tHost:", u.Host)
		fmt.Println("\tHostname:", u.Hostname())
		fmt.Println("\tOpaque:", u.Opaque)
		fmt.Println("\tPath:", u.Path)
		fmt.Println("\tRawPath:", u.RawPath)
		fmt.Println("\tRequestURI:", u.RequestURI())
		fmt.Println("\tQuery:", u.RawQuery)
		fmt.Println("\tString():", u.String())
		fmt.Println()
	} else {
		fmt.Printf("\tError: %v\n", err)
	}
}

func main() {
	// "absolute" URL
	str := "http://validURL.com/page"
	fmt.Println("Valid URL:", str)
	printURL(str)

	// url lib assumes we're working with a path here
	str = "Not a valid URL!!"
	fmt.Println("Invalid URL:", str)
	printURL(str)

	// We're trying hard to break url parser here
	str = "!@#($*)%&{][}"
	fmt.Println("Random symbol chars:", str)
	printURL(str)

	fmt.Println()

	str = "http://someplace.com?real=true&place=yes"
	fmt.Println("JoinPath test:", str)
	// Keeping things simple here, so throwing away the error
	str, _ = url.JoinPath(str, "city", "Tokyo")
	str, _ = url.JoinPath(str, "building", "megacorp")
	printURL(str)
}
