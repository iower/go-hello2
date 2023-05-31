package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "store.db")
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
