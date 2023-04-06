package main

import "fmt"

func main() {
	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4))
	fmt.Println(f(6, 7))

	var ff func(int, int) int = func(x, y int) int { return x + y }

	fmt.Println(ff(3, 4))
	fmt.Println(ff(6, 7))

	//
	action(10, 25, func(x int, y int) int { return x + y })
	action(5, 6, func(x int, y int) int { return x * y })

	fff := selectFn(1)
	fmt.Println(fff(2, 3))
	fmt.Println(fff(4, 5))
	fmt.Println(fff(7, 6))

	// closure
	fClosure := square()
	fmt.Println(fClosure())
	fmt.Println(fClosure())
	fmt.Println(fClosure())
}

func action(n1, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x, y int) int { return x + y }
	} else if n == 2 {
		return func(x, y int) int { return x - y }
	} else {
		return func(x, y int) int { return x * y }
	}
}

func square() func() int {
	x := 2
	return func() int {
		x++
		return x * x
	}
}
