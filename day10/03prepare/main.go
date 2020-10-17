package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gofullstack"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open database failed:", err)
		return
	}
}

func perpareInsert(name string, age uint8) {
	sqlStr := "INSERT INTO user(name,age) VALUES (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed:", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, age)
	if err != nil {
		fmt.Println("exec failed:", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get LastInsertId failed;", err)
		return
	}
	fmt.Println("inserted id:", id)
}

func main() {
	defer db.Close()
	perpareInsert("广东最靓的仔", 99)
}
