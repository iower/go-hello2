package main

import "fmt"

func main() {
	var numbers [5]int
	var numbers2 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	numbers3 := [4]int{1, 2, 3, 4}

	numbers4 := [3]int{1, 2}

	var numbers5 = [...]int{1, 2, 3, 4, 5}

	fmt.Println(numbers, numbers2, numbers3, numbers4, numbers5)

	fmt.Println(numbers5[0], numbers5[4])

	numbers5[0] = 11
	numbers5[4] = 55
	fmt.Println(numbers5[0], numbers5[4])

	colors := [3]string{2: "blue", 0: "red", 1: "greed"}
	fmt.Println(colors)

}
