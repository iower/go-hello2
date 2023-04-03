package main

import "fmt"

func main() {
	var f func(int, int) int = add
	fmt.Println(f(3, 4))

	var x = f(4, 5)
	fmt.Println((x))
}

func add(x, y int) int {
	return x + y
}
