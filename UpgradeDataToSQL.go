package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("Connection success.")
	}
	defer db.Close()
	
	Upgrading, err := db.Exec("UPDATE employee SET first_name = 'C', last_name = 'CC' WHERE id = 3")
	if err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("Updating success.")
	}
	count, err := Upgrading.RowsAffected()
	if err != nil{
		panic (err)
	}
	fmt.Println(count)
	defer db.Close()
}