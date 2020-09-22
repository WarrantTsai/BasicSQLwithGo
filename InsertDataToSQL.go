package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql","root:0000@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println("Connection success.")
	}
	defer db.Close()
	
	insert, err := db.Query("INSERT INTO employee(id, first_name, last_name) VALUES (2, 'A', 'AA')")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println("Insert new data success.")
	}
	
	defer insert.Close()
	
}