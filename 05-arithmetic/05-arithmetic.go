package main

import "fmt"

func main() {
	a := 4
	b := 6

	aa := 4.0
	bb := 6.0

	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := aa / bb

	h := 7 % 3

	a++
	b--

	fmt.Println(c, d, e, f, g, h, a, b)
}
