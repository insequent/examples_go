package main

import (
	"flag"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	url := flag.String("u", "ws://localhost", "The destination URL to initialize a websocket connection to")
	flag.Parse()

	log.Println("Starting connection to", *url)
	// TODO: origin?
	ws, err := websocket.Dial(*url, "", "http://localhost/")
	if err != nil {
		log.Fatalln("Failed to start websocket connection:", err)
	}
	defer ws.Close()

	log.Println("Starting listening loop...")
	for {
		msg := []byte{}
		if _, err := ws.Read(msg); err != nil {
			log.Fatalln("Failed to retrieve message from websocket")
		}
		log.Println("Received:", string(msg))

		log.Println("Sleeping (10s)...")
		time.Sleep(10 * time.Second)
	}
}
