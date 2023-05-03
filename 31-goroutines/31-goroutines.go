package main

import "fmt"

func main() {
	for i := 1; i < 7; i++ {
		go factorial(i)
	}
	fmt.Scanln()
	fmt.Println("The end")
}

func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
}
