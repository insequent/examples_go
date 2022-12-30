//         Simple HTTP Server

package main

import (
	"io"
	"log"
	"net/http"
)

func Print(w http.ResponseWriter, req *http.Request) {
	log.Println("Request")
	for key, value := range req.Header {
		log.Printf("\t%s: %s\n", key, value)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer req.Body.Close()
	log.Printf("\tBody:\n%s", string(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Print)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error on server stop:", err)
	}
}
