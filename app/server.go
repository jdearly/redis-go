package main

import (
	"fmt"
	"log"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go pingCommand(conn)
	}
}

func pingCommand(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}
	args := strings.Split(string(buffer[:]), "\n")
	responseStr := fmt.Sprintf("+PONG\r\n")
	for i := 0; i < len(args); i++ {
		// write data to response
		conn.Write([]byte(responseStr))
	}
}
