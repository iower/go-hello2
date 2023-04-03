package main

import "fmt"

func main() {
	hello()
	hello()
	hello()

	add(4, 5)
	add(20, 6)

	add2(1, 2, 3.4, 4.5, 5.6)

	// arguments - by value
	var a = 1
	fmt.Println("a before: ", a)
	inc(a)
	fmt.Println("a after: ", a)

	//
	addFew(1, 2, 3)
	addFew(1, 2, 3, 4)
	addFew(1, 2, 3, 4, 5)

	var nums = []int{5, 6, 7}
	addFew(nums...)

	// return
	fmt.Println(addReturn(4, 5))
	fmt.Println(addNamedReturn(4, 5))

	// multiple return
	age, name := multipleReturn(4, 5, "Tom", "Simpson")
	fmt.Println(age, name)
}

func hello() {
	fmt.Println("Hello!")
}

func add(x, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)
}

func add2(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func inc(x int) {
	fmt.Println("x before: ", x)
	x++
	fmt.Println("x after: ", x)
}

func addFew(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

func addReturn(x, y int) int {
	return x + y
}

func addNamedReturn(x, y int) (z int) {
	z = x + y
	return
}

func multipleReturn(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + "" + lastName
	return z, fullName
}
