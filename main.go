package main

import (
	"fmt"
	"net"
)

//type client interface {
//	message() string
//	name() string
//}

func main() {
	// creating tcp connection
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Print("server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// buffer to read input

	buffer := make([]byte, 1024)
	for {
		readBuffer, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading buffer: ", err)
			return
		}
		fmt.Printf("Server: %s\n", buffer[:readBuffer])

		data := []byte("client:")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Print("Error writing in server: ", err)
			return
		}
	}
}
