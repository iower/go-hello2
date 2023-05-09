package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 1
	intCh <- 2
	close(intCh)
	// intCh <- 3 // panic: send on closed channel
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)

	// default value 0
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)

	val, opened := <-intCh
	fmt.Println(val, opened)

	//
	intCh2 := make(chan int, 3)
	intCh2 <- 1
	intCh2 <- 2
	close(intCh2)

	for i := 0; i < cap(intCh2); i++ {
		if val, opened := <-intCh2; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Closed!")
		}
	}

}
