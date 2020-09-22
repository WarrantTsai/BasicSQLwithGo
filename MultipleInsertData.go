package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"strings"
)

func main(){
	db, err := sql.Open("mysql","root:0000@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println("Connection success.")
	}
	defer db.Close()
	
	var data = []map[string]string{
		{"v1":"1","v2":"A","v3":"AA"},
		{"v1":"2","v2":"B","v3":"BB"},
		{"v1":"3","v2":"C","v3":"CC"},
		{"v1":"4","v2":"D","v3":"DD"},
		{"v1":"5","v2":"E","v3":"EE"},
		{"v1":"6","v2":"F","v3":"FF"},
		{"v1":"7","v2":"G","v3":"GG"},
		{"v1":"8","v2":"H","v3":"HH"},
		{"v1":"9","v2":"I","v3":"II"},
		{"v1":"10","v2":"J","v3":"JJ"},
	}

	
	var sqlStr string = "INSERT INTO employee(id, first_name, last_name) VALUES "
	for _, row := range data{
		var s string = "('" + row["v1"] + "', '" + row["v2"] + "', '" + row["v3"] + "'),"
		sqlStr += s
	}
	// Trim the last word ',', 2 ways below
	//sqlStr = sqlStr[0:len(sqlStr)-1]
	sqlStr = strings.TrimSuffix(sqlStr, ",")
	fmt.Println(sqlStr)
	stmt, _ := db.Exec(sqlStr)
	fmt.Println(stmt)
}