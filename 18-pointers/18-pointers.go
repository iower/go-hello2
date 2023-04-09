package main

import "fmt"

func main() {
	var p0 *int
	fmt.Println(p0)

	// get address
	// fmt.Println(*p0) // err

	var x int = 4
	var p *int
	p = &x
	// var p *int = &x

	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)

	// change value via pointer
	*p = 25
	fmt.Println("Updated x", x)

	// shorted definition
	f := 2.3
	pf := &f
	fmt.Println("Address: ", pf)
	fmt.Println("Value: ", *pf)
	fmt.Println("Value: ", f)

	// nil pointer
	var pfEmpty *float64
	if pfEmpty != nil {
		fmt.Println("Value:", *pfEmpty)
	} else {
		fmt.Println("nil pointer")
	}

	// new - returns a pointer
	pNew := new(int)
	fmt.Println("New addr:", pNew)
	fmt.Println("New value:", *pNew)

	*pNew = 8
	fmt.Println("New value updated:", *pNew)
}
