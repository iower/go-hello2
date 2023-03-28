package main

import "fmt"

func main() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}

	switch a {
	case 6:
		fmt.Println("a == 6")
	case 7:
		fmt.Println("a == 7")
	case 4, 5:
		fmt.Println("a == 4 or 5")
	default:
		fmt.Println("other")
	}
}
