package main

import (
	"fmt"
	"net"
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

		// receive
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Println("Translation:", string(buff[0:n]))
	}
}
