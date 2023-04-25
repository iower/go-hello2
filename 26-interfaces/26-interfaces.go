package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
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

	tesla2 := Car{}
	boing2 := Aircraft{}
	drive(tesla2)
	drive(boing2)
}
