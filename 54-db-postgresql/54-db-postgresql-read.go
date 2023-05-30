package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	conn := "user=postgres password=password dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}
