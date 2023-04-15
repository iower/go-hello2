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
}
