package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	socketPath := "echo.sock"
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to initialize socket '%s': %v", socketPath, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer os.Remove(socketPath)
	defer cancel()

	// Catch signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		case <-ch:
			os.Remove(socketPath)
			os.Exit(1)
		}
	}(ctx)

	// Main loop
	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatalf("Failed to accept next connection: %v", err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			data := strings.Builder{}
			b := make([]byte, 1)
			for {
				_, err := conn.Read(b)
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}

					log.Fatalf("Failed to read from connection: %v", err)
				}

				if b[0] == '\n' {
					log.Print("RECV: ", string(data.String()))
					data.Reset()
				} else {
					data.WriteByte(b[0])
				}
			}

		}(conn)
	}
}
