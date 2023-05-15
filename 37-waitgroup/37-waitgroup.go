package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Goroutine %d stared\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutine %d finished\n", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
	fmt.Println("All finished")
}
