package main

import (
	"fmt"
	"io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

type phoneWriter struct{}

func (p phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}

func main() {
	// Read
	phone1 := phoneReader("+1(234)567 8900")

	buffer := make([]byte, len(phone1))
	fmt.Println(buffer)

	phone1.Read(buffer)
	fmt.Println(buffer)
	fmt.Println(string(buffer))

	phone2 := phoneReader("+2-345-678-99-99")
	buffer = make([]byte, len(phone2))
	fmt.Println(buffer)

	phone2.Read(buffer)
	fmt.Println(buffer)
	fmt.Println(string(buffer))

	// Write
	fmt.Println()

	bytes1 := []byte("+1(234)567 8900")
	fmt.Println(bytes1)
	writer := phoneWriter{}
	writer.Write(bytes1)

	bytes2 := []byte("+2-345-678-99-99")
	writer.Write(bytes2)
}
