/*
CREATE TABLE products(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  model TEXT,
  company TEXT,
  price INTEGER
);
*/

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
	result, err := db.Exec("insert into products (model, company, price) values ('iPhone X', $1, $2)", "Apple", 1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
