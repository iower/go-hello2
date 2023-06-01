package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id      bson.ObjectId `bson: "_id"`
	Model   string        `bson: "model"`
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
	query := bson.M{}
	products := []Product{}
	productCollection.Find(query).All(&products)

	for _, p := range products {
		fmt.Println(p.Model, p.Company, p.Price)
	}
}
