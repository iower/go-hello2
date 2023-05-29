package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
create database productdb;
use productdb;
create table products (
    id int auto_increment primary key,
    model varchar(30) not null,
    company varchar(30) not null,
    price int not null
)
*/

func main() {
	db, err := sql.Open("mysql", "root:root@/productdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into productdb.products (model, company, price) values (?, ?, ?)", "iPhone X", "Apple", 1000)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
