package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id      bson.ObjectId `bson:"_id"`
	Model   string        `bson:"model"`
	Company string        `bson: "company"`
	Price   int           `bson: "price"`
}

func main() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	productCollection := session.DB("productdb").C("products")

	p1 := &Product{
		Id:      bson.NewObjectId(),
		Model:   "iPhone 8",
		Company: "Apple",
		Price:   1000,
	}
	err = productCollection.Insert(p1)
	if err != nil {
		fmt.Println(err)
	}

	p2 := &Product{
		Id:      bson.NewObjectId(),
		Model:   "Pixel 2",
		Company: "Google",
		Price:   800,
	}
	p3 := &Product{
		Id:      bson.NewObjectId(),
		Model:   "Xplay7",
		Company: "Vivo",
		Price:   700,
	}

	err = productCollection.Insert(p2, p3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1, p2, p3)
}
