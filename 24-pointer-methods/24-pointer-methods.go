package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) updateAge_(newAge int) {
	p.age = newAge
}

func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}

func main() {
	tom := person{name: "Tom", age: 24}
	fmt.Println("before", tom.age)
	tom.updateAge_(25)
	fmt.Println("after", tom.age)

	tom.updateAge(25)
	fmt.Println("after2", tom.age)

	tomPointer := &tom
	tomPointer.updateAge(26)
	fmt.Println("after3", tom.age)
}
