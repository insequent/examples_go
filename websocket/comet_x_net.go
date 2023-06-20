package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

func Read(wc *websocket.Conn) {
	log.Println("Starting listening loop...")

	for {
		var msg string
		if err := websocket.Message.Receive(wc, &msg); err != nil {
			log.Fatalln("Error while reading from websocket:", err)
		}
		msg = strings.TrimSpace(msg)
		if len(msg) > 0 {
			log.Printf("Recieved: %d\n", len(msg))
			return
		}
		time.Sleep(time.Millisecond)
	}
}

func main() {
	url := flag.String("url", "ws://127.0.0.1:8060/api/v1/events/stream", "The destination URL to initialize a websocket connection to")
	username := flag.String("u", "admin", "Username used for authentication")
	password := flag.String("p", "admin", "Password used for authentication")
	flag.Parse()

	log.Println("Starting connection to", *url)
	wc, err := websocket.Dial(*url, "", "http://localhost/")
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	defer wc.Close()

	// Listen for events
	Read(wc)

	// Authentication
	log.Println("Sending authentication request")
	for _, msg := range []string{*username, "Password", *password, "", ""} {
		if err := websocket.Message.Send(wc, msg); err != nil {
			log.Fatalln("Failed to send message:", msg, err)
		}
	}

	// A temporary solution
	time.Sleep(10 * time.Minute)
}
