package main

import (
	"fmt"
	"os"
)

func exitOnErr(err error) {
	if err != nil {
		fmt.Println("Unable to use file:", err)
		os.Exit(1)
	}
}

func main() {
	file, err := os.Create("hello.txt")
	fmt.Println(file.Name(), file, err)
	exitOnErr(err)
	defer file.Close()

	file2, err2 := os.Open("hello.txt")
	fmt.Println(file2.Name(), file2, err2)
	exitOnErr(err2)
	defer file2.Close()

	//
	file3, err3 := os.OpenFile("hi.txt", os.O_RDONLY, 0666)
	fmt.Println(file3.Name(), file3, err3)
	exitOnErr(err3)
	defer file3.Close()

	file4, err4 := os.OpenFile("hi2.txt", os.O_WRONLY, 0666)
	fmt.Println(file4.Name(), file4, err4)
	exitOnErr(err4)
	defer file4.Close()
}
