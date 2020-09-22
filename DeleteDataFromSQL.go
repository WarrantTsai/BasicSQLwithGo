package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println("Connection success.")
	}
	defer db.Close()

	delete, err := db.Query("DELETE FROM employee WHERE id = 2 AND first_name = 'A' AND last_name = 'AA'")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println(delete)
		fmt.Println("Data deleted.")
	}
}