package main

import "fmt"

func main() {
	defer finish()
	fmt.Println("Started")
	fmt.Println("Working")
	fmt.Println(divide(15, 5))
	defer fmt.Println("Working2")
	fmt.Println(divide(4, 0))
	defer fmt.Println("Working3")
}

func finish() {
	fmt.Println("Finished")
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Divistion by zero!")
	}
	return x / y
}
