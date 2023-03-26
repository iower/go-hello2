package main

import (
	"fmt"
	"reflect"
)

func main() {
	const pi float32 = 3.1415
	// pi = "immutable"

	fmt.Println(pi)

	const (
		pi2 float64 = 3.1415
		e2  float64 = 2.7182
	)

	const pi3, e3 = 3.1415, 2.7182

	const (
		a = 1
		b
		c
		d = 2
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)

	const n = 5
	fmt.Println(n, reflect.TypeOf(n))

	// var mvar = 7
	// const mconst = mvar // cannot init by var

	const nn = n
}
