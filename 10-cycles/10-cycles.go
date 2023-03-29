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

	// range
	users := [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}
	for _, value := range users {
		fmt.Println(value)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(i, users[i])
	}

	// break, continue
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	for _, value := range numbers {
		if value < 0 {
			continue
		}
		sum += value
	}
	fmt.Println("Sum:", sum)

	var numbers2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum2 = 0

	for _, value := range numbers2 {
		if value > 4 {
			break
		}
		sum2 += value
	}
	fmt.Println("Sum2:", sum2)
}
