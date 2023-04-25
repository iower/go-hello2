package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{}

func (c Car) move() {
	fmt.Println("car move")
}

type Aircraft struct{}

func (a Aircraft) move() {
	fmt.Print("aircraft move")
}

func main() {
	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()
}
