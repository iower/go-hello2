package main

import "fmt"

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibbonachi(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))

	fmt.Println(fibbonachi(4))
	fmt.Println(fibbonachi(5))
	fmt.Println(fibbonachi(6))
}
