package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/productdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update productdb.products set price = ? where id = ?", 900, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
