package main

import "fmt"

func main() {
	fmt.Println(2 >> 1)
	fmt.Println(2 << 1)

	fmt.Println(4 >> 1)
	fmt.Println(4 << 2)

	fmt.Println("===")

	fmt.Println(1 & 1)
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)

	fmt.Println(0 ^ 1)

	fmt.Println(1 &^ 0)
}
