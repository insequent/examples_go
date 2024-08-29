//         Simple HTTP Server

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Stream(w http.ResponseWriter, r *http.Request) {
	log.Println("Request")
	for key, value := range r.Header {
		log.Printf("\t%s: %s\n", key, value)
	}
	w.Header().Add("sessionID", fmt.Sprintf("%d", rand.Int()))
	w.WriteHeader(200)

	// Write triggers http.OK header
	flusher, _ := w.(http.Flusher)
	w.Write([]byte("{\"message\":\"first\"}:::"))
	flusher.Flush()
	time.Sleep(10 * time.Second)
	w.Write([]byte("{\"message\":\"second\"}:::"))
	flusher.Flush()
	time.Sleep(10 * time.Second)
	w.Write([]byte("{\"message\":\"third\"}:::"))
	flusher.Flush()
	time.Sleep(10 * time.Second)
	w.Write([]byte("{\"message\":\"fourth\"}"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Stream)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error on server stop:", err)
	}
}
