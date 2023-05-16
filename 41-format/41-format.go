package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprint(file, "It's a good ")
	fmt.Fprintln(file, "weather today!")
}
