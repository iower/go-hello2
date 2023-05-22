package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter name: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Enter age: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Println(name, age)

	// Scan, Scanln
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Println(name, age)

	//
	fmt.Print("Enter name (space) age: ")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)
}
