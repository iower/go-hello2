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

	//
	slice := []int{-2, 4, 3, -1, 7, -4, 23}
	sumOfEvens := sum(slice, isEven)
	fmt.Println(sumOfEvens)
	sumOfPositives := sum(slice, isPositive)
	fmt.Println(sumOfPositives)

	// return function
	f = selectFn(1)
	fmt.Println(f(3, 4))
	f = selectFn(3)
	fmt.Println(f(3, 4))
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
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

func isEven(n int) bool {
	return n%2 == 0
}

func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {
	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}
