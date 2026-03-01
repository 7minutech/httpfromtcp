package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const READ_SIZE = 8

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("failed to open message file: %v", err)
	}
	defer file.Close()

	var buffer []byte = make([]byte, READ_SIZE)
	var line strings.Builder
	for {
		n, err := file.Read(buffer)
		if errors.Is(err, io.EOF) {
			fmt.Printf("read: %s\n", line.String())
			os.Exit(0)
		}
		if err != nil {
			log.Fatalf("failed to read from message file: %v", err)
		}
		msg := string(buffer[:n])
		parts := strings.Split(msg, "\n")
		if len(parts) > 1 {
			for i := 0; i < len(parts)-1; i++ {
				line.WriteString(parts[i])
				fmt.Printf("read: %s\n", line.String())
				line.Reset()
			}
			line.WriteString(parts[len(parts)-1])

		} else {
			line.WriteString(msg)
		}
	}
}
