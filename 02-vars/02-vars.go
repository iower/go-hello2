package main

import "fmt"

func main() {
	var hello string
	var a, A, b, c int

	hello = "Hello world"
	var hello2 string = "Hello world 2"

	a = 1
	A = 11
	b = 2
	c = 3

	// bulk
	var (
		d int = 4
		e int = 5
	)

	var userName, userAge = "Tom", 25

	fmt.Println(hello, hello2, a, A, b, c, d, e, userName, userAge)

	// change var
	hello = "New hello"
	fmt.Println(hello)

	// short
	name := "Alice"
	fmt.Println(name)
}
