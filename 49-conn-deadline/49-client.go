package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		fmt.Print("Enter a word: ")
		var source string
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Wrong input", err)
			continue
		}

		// send
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// get response
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Translation:", string(buff[0:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
	}
}
