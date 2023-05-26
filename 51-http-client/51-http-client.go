package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	time.Sleep(2 * time.Second)
	fmt.Println("=================")

	// Request
	req, err := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "MSIE/15.0")

	client2 := &http.Client{}
	resp2, err2 := client2.Do(req)
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp2.Body)
}
