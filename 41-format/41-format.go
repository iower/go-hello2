package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprint(file, "It's a good ")
	fmt.Fprintln(file, "weather today!")

	//
	file2, err2 := os.Create("formatted.txt")
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	defer file2.Close()
	fmt.Fprintf(file2, "boolean: %t %t \n", false, true)
	fmt.Fprintf(file2, "binary: %b %b \n", 4, 5)
	fmt.Fprintf(file2, "char: %c \n", 1234)
	fmt.Fprintf(file2, "decimal: %d \n", 10)
	fmt.Fprintf(file2, "octal: %o \n", 9)
	fmt.Fprintf(file2, "quotes: %q %q \n", '#', "quoted")
	fmt.Fprintf(file2, "hex: %x %X \n", 255, 255)
	fmt.Fprintf(file2, "unicode: %U \n", 1)
	fmt.Fprintf(file2, "exponential: %e %E \n", 1000.0, 1000.0)
	fmt.Fprintf(file2, "float: %f %F \n", 123.456, 123.456)
	fmt.Fprintf(file2, "float|exp: %g %g %G %G \n", 123.456, 123.456, 123123123.456, 123123123.456)
	fmt.Fprintf(file2, "string: %s \n", "string")
	v := 123
	fmt.Fprintf(file2, "pointer: %p \n", &v)

	fmt.Fprintf(file2, "value: %v %v %v %v", true, 10, 1.23, "string")
}
