package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Add("accept", "application/json")
	if err != nil {
		log.Printf("Error creating request: %v", err)
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making http request: %v", err)
		os.Exit(1)
	}
	for k, v := range resp.Header {
		log.Printf("Header :: %s: %v\n", k, v)
	}

	result := []byte{}
	defer resp.Body.Close()
	for {
		p := make([]byte, 1)
		n, err := resp.Body.Read(p)
		switch {
		case err != nil && errors.Is(err, io.EOF):
			// All done!
			fmt.Println("JSON Received:")
			if n == 1 {
				result = append(result, p[0])
			}
			fmt.Println(string(result))
			return
		case err != nil:
			log.Panicf("Failed to read single byte from connection: %v", err)
			return
		case n <= 0:
			continue
		case p[0] == ':' && len(result) > 2 && result[len(result)-1] == ':' && result[len(result)-2] == ':':
			fmt.Println("JSON Received:")
			fmt.Println(string(result[:len(result)-2]))
			result = []byte{}
		default:
			result = append(result, p[0])
		}
	}
}
