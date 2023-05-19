package main

import (
	"fmt"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {
	fmt.Fprintln(os.Stdout, "Hello stdout")
	fmt.Fprintf(os.Stdout, "%.2f", 123456789.123456789)

	fmt.Println()
	fmt.Println("1", 11)
	fmt.Print("2")
	fmt.Printf("%d", 3)

	fmt.Println()

	tom := person{
		name:   "Tom",
		age:    24,
		weight: 68.5,
	}
	fmt.Printf("%-10s %-10d %-10.3f\n", tom.name, tom.age, tom.weight)
}
