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
	
	rows, err := db.Query("SELECT first_name, last_name FROM employee WHERE id != 2")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var firstname, lastname string
		err = rows.Scan(&firstname, &lastname)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(firstname, lastname,)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}