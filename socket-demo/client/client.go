package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to server at localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	
	fmt.Println("Connected to server. Type messages and press enter to send.")

	// Read server welcome message
	reader := bufio.NewReader(conn)
	welcome, _ := reader.ReadString('\n')
	fmt.Print("Server says:", welcome)

	// Read from stdin and send to server
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			conn.Write([]byte(text + "\n"))
		}
	}()

	// Print responses from server
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by server")
			return
		}
		fmt.Print("Server response:", response)
	}
}