package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const NETWORK = "udp"
const ADDRESS = "localhost:42069"

func main() {
	addr, err := net.ResolveUDPAddr(NETWORK, ADDRESS)
	if err != nil {
		log.Printf("failed to resolve ip address: %s, %v", ADDRESS, err)
	}
	conn, err := net.DialUDP(NETWORK, nil, addr)
	if err != nil {
		log.Printf("failed to dial udp: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error: %v", err)
		}
		conn.Write([]byte(line))

	}

}
