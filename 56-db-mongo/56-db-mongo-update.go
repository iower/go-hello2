package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id      bson.ObjectId `bson:"_d"`
	Model   string        `bson:"model"`
	Company string        `bson:"company"`
	Price   int           `bson:"price"`
}

func main() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	productCollection := session.DB("productdb").C("products")

	err = productCollection.Update(bson.M{"model": "iPhone 8"}, bson.M{"$set": bson.M{"price": 1100}})
	if err != nil {
		fmt.Println(err)
	}

	products := []Product{}
	productCollection.Find(bson.M{}).All(&products)

	for _, p := range products {
		fmt.Println(p.Model, p.Company, p.Price)
	}
}
