package main

import "fmt"

func main() {
	people := map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people)
	fmt.Println(people["Alice"])
	fmt.Println(people["Test"])
	fmt.Println(people["Test2"])

	people["Alice"] = 88
	fmt.Println(people["Alice"])

	// addition
	people["Test"] = 100
	fmt.Println(people)

	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}

	// missing keys
	_, ok := people["Unknown"]
	fmt.Println(ok)

	if val, ok := people["Unknown"]; ok {
		fmt.Println(val)
	}

	// for range
	for key, value := range people {
		fmt.Println(key, value)
	}

	// make empty
	peopleEmpty := make(map[string]int)
	fmt.Println(peopleEmpty)

	// delete
	fmt.Println(people)
	delete(people, "Test")
	fmt.Println(people)

	delete(people, "Missing")
	delete(people, "Missing2")
}
