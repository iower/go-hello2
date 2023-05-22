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

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func writeData() {
	tom := person{
		name:   "Tom",
		age:    24,
		weight: 68.5,
	}

	file := createFile("data.txt")
	defer file.Close()

	fmt.Fprintln(file, tom.name)
	fmt.Fprintln(file, tom.age)

	//
	file2 := createFile("data2.txt")
	defer file2.Close()
	fmt.Fprintf(file2, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func readData() {
	var name string
	var age int
	var weight float64

	file := openFile("data.txt")
	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)

	//
	file2 := openFile("data2.txt")
	defer file2.Close()

	_, err := fmt.Fscanf(file2, "%s %d %f\n", &name, &age, &weight)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
}

func main() {
	writeData()
	readData()
}
