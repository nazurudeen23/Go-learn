package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Start TCP server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080")

	for {
		// Wait for client connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}

		// Handle client in a new goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Get client address
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected:", clientAddr)

	// Send welcome message
	conn.Write([]byte("Welcome to the server!\n"))

	// Read data from client
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n') // Corrected redeclaration of message
		if err != nil {
			fmt.Println("Client disconnected:", clientAddr)
			return
		}
		
		// Print and echo back the message
		fmt.Printf("Received from %s: %s", clientAddr, message)
		conn.Write([]byte("Echo: " + message))
	}
}