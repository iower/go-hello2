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

	// pointers
	var tomPointer *person = &tom
	fmt.Println("tomPointer", tomPointer)
	tomPointer.age = 29
	fmt.Println(tom.age)
	(*tomPointer).age = 30
	fmt.Println(tom.age)

	var tom2 *person = &person{name: "Tom2", age: 23}
	var anon *person = new(person)
	fmt.Println(tom2, anon)

	// fileld pointers
	var agePointer *int = &tom.age
	*agePointer = 35
	fmt.Println(tom)
}
