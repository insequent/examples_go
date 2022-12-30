//         Simple HTTP Server

package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	conns := make(map[net.Conn]http.ConnState)

	s.ConnState = func(c net.Conn, cs http.ConnState) {
		switch cs {
		case http.StateNew:
			conns[c] = http.StateNew
		case http.StateActive:
			conns[c] = http.StateActive
		case http.StateIdle:
			conns[c] = http.StateIdle
		case http.StateHijacked:
			conns[c] = http.StateHijacked
		case http.StateClosed:
			conns[c] = http.StateClosed
		}
	}

	http.HandleFunc("/", hello)
	fmt.Println("Starting child http process...")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Serve(ln)
	}()

	fmt.Println("Starting 10sec timer")
	ticker := time.NewTicker(time.Second * time.Duration(10))

	fmt.Println("Starting main for loop")

loop:
	for {
		select {
		case <-ticker.C:
			fmt.Println("Times up! See ya")
			fmt.Printf("Connections? %s\n", conns)
			ln.Close()
			break loop
		}
	}

	fmt.Println("Waiting on server to finish...")
	wg.Wait()
}
