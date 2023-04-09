package main

import "fmt"

func main() {
	d := 5
	fmt.Println("d before:", d)
	changeValue_test(d)
	fmt.Println("d after (not changed):", d)

	changeValue(&d)
	fmt.Println("d after (changed):", d)

	// create pointer
	p1 := createPointer(7)
	p2 := createPointer(10)
	fmt.Println(p1, *p1)
	fmt.Println(p2, *p2)
}

func changeValue_test(x int) {
	x = x * x
}

func changeValue(p *int) {
	*p = (*p) * (*p)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}
