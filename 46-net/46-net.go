package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	httpRequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n"
	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 2)) // close HTTP/1.1

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
