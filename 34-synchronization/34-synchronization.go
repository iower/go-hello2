package main

import "fmt"

func main() {
	intCh := make(chan int)
	go factorial(5, intCh)
	fmt.Println(<-intCh)
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}
