package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}

	var i = 10
	for ; i >= 1; i-- {
		fmt.Println(i * i)
	}

	var j = 1
	for j <= 10 {
		fmt.Println(j * j)
		j++
	}

	for j >= 1 {
		fmt.Println(j * j)
		j--
	}

	// nested cycles
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	// ...
}
