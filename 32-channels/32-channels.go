package main

import "fmt"

func main() {
	var intCh chan int
	fmt.Println(intCh)

	intCh <- 5
	val := <-intCh
	fmt.Println(val)
}
