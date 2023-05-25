package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	message := "Hello, I am a server"

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("Connect!", conn.RemoteAddr())
		conn.Write([]byte(message))
		conn.Close()
	}
}
