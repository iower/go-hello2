package main

import "fmt"

type mile int
type kilometer int

type library []string

func distanceToEnemy(distance mile) {
	fmt.Println(distance, "miles")
}

func printBooks(lib library) {
	for _, value := range lib {
		fmt.Println(value)
	}
}

func main() {
	var distance mile = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)

	var distance2 kilometer = 5
	fmt.Println(distance2)

	distanceToEnemy(distance)

	var myLibrary library = library{"Book1", "Book2", "Book3"}
	printBooks(myLibrary)
}
