package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func Listen(conn *websocket.Conn) {
	log.Println("Starting listening loop...")
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatalln("Error while reading from websocket:", err)
		}
		log.Printf("Recieved:\n%s", message)
	}
}

func main() {
	url := flag.String("url", "ws://localhost", "The destination URL to initialize a websocket connection to")
	username := flag.String("u", "admin", "Username used for authentication")
	password := flag.String("p", "admin", "Password used for authentication")
	flag.Parse()

	log.Println("Starting connection to", *url)
	conn, _, err := websocket.DefaultDialer.Dial(*url, http.Header{"Origin": {"http://localhost/"}})
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	defer conn.Close()

	// Listen for events
	go Listen(conn)

	// Authentication
	log.Println("Sending authentication request")
	for _, msg := range []string{*username, "Password", *password, "", ""} {
		if err = conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Fatalln("Failed to send message:", msg, err)
		}
	}

	// A temporary solution
	time.Sleep(10 * time.Minute)
}
