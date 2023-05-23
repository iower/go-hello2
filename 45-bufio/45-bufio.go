package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}

	file, err := os.Create("some.dat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}

	writer.Flush()

	//
	file2, err := os.Open("some.dat")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file2.Close()

	reader := bufio.NewReader(file2)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}
