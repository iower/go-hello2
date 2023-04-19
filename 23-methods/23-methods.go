package main

import "fmt"

type library []string

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
}
func (p person) eat(meal string) {
	fmt.Println(p.name, "eats", meal)
}

func main() {
	myLibrary := library{"book1", "book2", "book3"}
	myLibrary.print()

	tom := person{name: "Tom", age: 24}
	tom.print()
	tom.eat("Soup")
}
