package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name        string
	age         int
	contactInfo contact
}

type person2 struct {
	name string
	age  int
	contact
}

/* // recursive types are not allowed
type node struct{
	value int
	next node
}
*/

type node struct {
	value int
	next  *node
}

func printNodeValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func main() {
	var tom = person{
		name: "Tom",
		age:  24,
		contactInfo: contact{
			email: "tom@server.com",
			phone: "+1234567890",
		},
	}

	fmt.Println(tom)
	tom.contactInfo.email = "supertom@server.com"
	fmt.Println(tom)

	var tom2 = person2{
		name: "Tom2",
		age:  42,
		contact: contact{
			email: "tom2@server.com",
			phone: "+2222222222",
		},
	}
	fmt.Println(tom2)

	//
	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
