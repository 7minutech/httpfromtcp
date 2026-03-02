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

	ch := getLinesChannel(file)

	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	var buffer []byte = make([]byte, READ_SIZE)
	var line strings.Builder

	go func() {
		defer close(ch)
		defer f.Close()
		for {
			n, err := f.Read(buffer)
			if errors.Is(err, io.EOF) {
				ch <- line.String()
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
