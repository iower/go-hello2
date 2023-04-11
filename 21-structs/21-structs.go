package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{"Tom", 23}
	var mot person = person{"Mot", 32}
	var alice person = person{age: 23, name: "Alice"}
	var bob = person{name: "Bob", age: 31}

	unknown := person{} // default values

	fmt.Println(tom, mot, alice, bob, unknown)

	// fields
	fmt.Println(tom.name, tom.age)
	tom.age = 38
	fmt.Println(tom.name, tom.age)
}
