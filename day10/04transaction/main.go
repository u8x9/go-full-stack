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

func transaction() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("begin transaction failed:", err)
		return
	}
	sqlStr := "UPDATE user SET age=age-3 WHERE id=?"
	stmt1, err := tx.Prepare(sqlStr)
	if err != nil {
		tx.Rollback()
		fmt.Println("Prepare 1  failed:", err)
		return
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(1)
	if err != nil {
		tx.Rollback()
		fmt.Println("Exec 1 failed:", err)
		return
	}
	sqlStr = "UPDATE user SET age=age+3 WHERE id=?"
	stmt2, err := tx.Prepare(sqlStr)
	if err != nil {
		tx.Rollback()
		fmt.Println("Prepare 2  failed:", err)
		return
	}
	defer stmt2.Close()
	_, err = stmt2.Exec(2)
	if err != nil {
		tx.Rollback()
		fmt.Println("Exec 2 failed:", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("Commit failed:", err)
		return
	}
	fmt.Println("Exec transaction success")
}

func main() {
	defer db.Close()
	transaction()
}
