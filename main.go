package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const READ_SIZE = 8

func main() {
	const TCP_ADDR = "127.0.0.1"
	const PORT = "42069"
	const NETWORK = "tcp"
	listner, err := net.Listen(NETWORK, TCP_ADDR+":"+PORT)
	if err != nil {
		log.Fatalf("failed to create tcp listener: %v", err)
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Printf("failed find connection for listener: %v", err)
			continue
		}

		fmt.Println("Connection has been accepted")

		ch := getLinesChannel(conn)

		for line := range ch {
			fmt.Println(line)
		}

		fmt.Println("channel has been closed")

	}

}

func getLinesChannel(conn io.ReadCloser) <-chan string {
	ch := make(chan string)

	var buffer []byte = make([]byte, READ_SIZE)
	var line strings.Builder

	go func() {
		defer close(ch)
		defer conn.Close()
		for {
			n, err := conn.Read(buffer)
			if errors.Is(err, io.EOF) {
				if line.Len() > 0 {
					ch <- line.String()
				}
				return
			}
			if err != nil {
				return
			}
			msg := string(buffer[:n])
			parts := strings.Split(msg, "\n")
			if len(parts) > 1 {
				for i := 0; i < len(parts)-1; i++ {
					line.WriteString(parts[i])
					ch <- line.String()
					line.Reset()
				}
				line.WriteString(parts[len(parts)-1])

			} else {
				line.WriteString(msg)
			}
		}
	}()

	return ch

}
