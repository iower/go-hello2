package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/productdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Exec
	result, err := db.Exec("INSERT INTO Products (model, company, price) VALUES ('iPhone X', 'Apple', '1000')")
	fmt.Println(result.LastInsertId())

	// Query
	rows, err2 := db.Query("SELECT name FROM users WHERE age=23")
	if err2 != nil {
		panic(err2)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

}
