package main

import "fmt"

type library []string

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func main() {
	myLibrary := library{"book1", "book2", "book3"}
	myLibrary.print()
}
