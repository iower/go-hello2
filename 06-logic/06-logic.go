package main

import "fmt"

func main() {
	a := 8
	b := 3

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	fmt.Println("===")
	fmt.Println(!false, !!false, !!!false)
	fmt.Println("===")

	c := false
	d := true
	fmt.Println(c && d, c || d)
}
