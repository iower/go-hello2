package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id      bson.ObjectId `bson:"_id"`
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

	_, err = productCollection.RemoveAll(bson.M{"company": "Vivo"})
	if err != nil {
		fmt.Println(err)
	}

	products := []Product{}
	productCollection.Find(bson.M{}).All(&products)

	for _, p := range products {
		fmt.Println(p.Model, p.Company, p.Price)
	}
}
