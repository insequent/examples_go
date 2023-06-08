//         Simple HTTP Server

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Print(w http.ResponseWriter, r *http.Request) {
	//log.Println("Request")
	//for key, value := range req.Header {
	//	log.Printf("\t%s: %s\n", key, value)
	//}

	blah := r.FormValue("test")
	if blah != "" {
		fmt.Println("Blah:", blah)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer r.Body.Close()
	//log.Printf("\tBody:\n%s", string(body))
	fmt.Println(string(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Print)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error on server stop:", err)
	}
}
