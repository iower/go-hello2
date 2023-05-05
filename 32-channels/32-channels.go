package main

import (
	"fmt"
	"time"
)

func main() {
	var intCh chan int = make(chan int)
	fmt.Println(intCh)

	go func() {
		fmt.Println("Go routine starts")
		intCh <- 5
		intCh <- 6
		fmt.Println("Go routine ends")
	}()

	fmt.Println("Wait...")
	fmt.Println(<-intCh)
	fmt.Println("The End")

	//
	fmt.Println()

	myCh := make(chan int)
	go factorial(5, myCh)
	fmt.Println("Wait2...")
	time.Sleep(2 * time.Second)
	fmt.Println(<-myCh)
	fmt.Println("The End 2")
}

func factorial(n int, ch chan int) {
	fmt.Println("Go factorial...")
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	fmt.Println(n, "-", res)

	fmt.Println("Send...")
	ch <- res
	fmt.Println("Go factorial ends")
}
