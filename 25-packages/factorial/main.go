package main

import (
	"fmt"

	"rsc.io/quote"

	"factorial/computation"
)

func main() {
	fmt.Println(computation.Factorial(4))
	fmt.Println(computation.Factorial(5))

	fmt.Printf(quote.Hello())
}
