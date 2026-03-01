package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const READ_SIZE = 8

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("failed to open message file: %v", err)
	}
	defer file.Close()

	var buffer []byte = make([]byte, READ_SIZE)
	for {
		n, err := file.Read(buffer)
		if errors.Is(err, io.EOF) {
			os.Exit(0)
		}
		if err != nil {
			log.Fatalf("failed to read from message file: %v", err)
		}
		fmt.Printf("read: %s\n", string(buffer[:n]))
	}
}
