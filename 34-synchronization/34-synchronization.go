package main

import "fmt"

func main() {
	results := make(map[int]int)
	structCh := make(chan struct{})

	go factorial(5, structCh, results)

	<-structCh
	for i, v := range results {
		fmt.Println(i, "-", v)
	}
}

func factorial(n int, ch chan struct{}, results map[int]int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
