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

	// buffered
	myChBuffered := make(chan int, 3)
	fmt.Println("cap", cap(myChBuffered), "len", len(myChBuffered))
	myChBuffered <- 1
	myChBuffered <- 2
	myChBuffered <- 3
	// myChBuffered <- 4 // blocking
	fmt.Println(<-myChBuffered)
	fmt.Println(<-myChBuffered)
	fmt.Println(<-myChBuffered)
	fmt.Scan()

	//
	fmt.Println("Start")
	fmt.Println(<-createChan(5))
	fmt.Println("End")
}

func factorial(n int, ch chan<- int) {
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

func createChan(n int) <-chan int {
	ch := make(chan int)
	go func() {
		ch <- n
	}()
	return ch
}
