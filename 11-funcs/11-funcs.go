package main

import "fmt"

func main() {
	hello()
	hello()
	hello()

	add(4, 5)
	add(20, 6)

	add2(1, 2, 3.4, 4.5, 5.6)

	// arguments - by value
	var a = 1
	fmt.Println("a before: ", a)
	inc(a)
	fmt.Println("a after: ", a)
}

func hello() {
	fmt.Println("Hello!")
}

func add(x, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)
}

func add2(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func inc(x int) {
	fmt.Println("x before: ", x)
	x++
	fmt.Println("x after: ", x)
}
