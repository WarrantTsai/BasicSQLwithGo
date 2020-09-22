package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Go MySQL!!!")
	
	db, err := sql.Open("mysql","root:0000@tcp(127.0.0.1:3306)/")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println("Successfully Connected to MySQL database")
	}
	defer db.Close()
	
	
	// CREATE DATABASE
	_,err = db.Exec("CREATE DATABASE testDB")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database.")
	}
	
	// SELECT DATABASE
	_,err = db.Exec("USE testDB")
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("DB selected successfully.")
	}
	
	// CREATE TABLE
	stmt, err := db.Prepare("CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY(id));")
	if err != nil{
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("Table created successfully.")
	}
	
	// CLOSE DATABASE
	defer db.Close()
	fmt.Println("Database closed")
}
