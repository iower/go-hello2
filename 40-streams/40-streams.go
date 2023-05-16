package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)

	fmt.Println()
	// custom reader

	phone1 := phoneReader("+1(234)567 89 99`")
	io.Copy(os.Stdout, phone1)
	fmt.Println()
}

type phoneReader string

func (p phoneReader) Read(bs []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}
