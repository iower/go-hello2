package main

import "fmt"

func main() {
	var f func(int, int) int = add
	fmt.Println(f(3, 4))
	f = multiply
	fmt.Println(f(3, 4))

	var x = f(4, 5)
	fmt.Println((x))

	var f1 func(string) = display
	f1("hello")

	//
	action(10, 15, add)
	action(5, 6, multiply)
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func display(message string) {
	fmt.Println(message)
}

func action(n1, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}
