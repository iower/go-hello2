package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	conn := "user=postgres password=password dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update products set price = $1 where id = $2", 1100, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
