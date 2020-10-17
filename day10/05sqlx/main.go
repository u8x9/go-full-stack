package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// User 用户
type User struct {
	Id   uint32
	Name string
	Age  uint8
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gofullstack"
	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open database failed:", err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Println("connection failed:", err)
		return
	}
}

// queryRow 查询单条记录
func queryRow() {
	sqlStr := "SELECT id,name,age FROM user WHERE id=?"
	var u User
	if err := db.Get(&u, sqlStr, 1); err != nil {
		fmt.Println("Get failed:", err)
		return
	}
	fmt.Println(u)
}

// query 查询多条记录
func query() {
	sqlStr := "SELECT id,name,age FROM user ORDER BY id DESC LIMIT ?"
	var us []User
	if err := db.Select(&us, sqlStr, 5); err != nil {
		fmt.Println("Select failed:", err)
		return
	}
	for _, u := range us {
		fmt.Println(u)
	}
}

// tx 事务
func tx() {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO user (name,age) VALUES(?,?)", "江西最帅的哥", 88)
	tx.MustExec("INSERT INTO user (name,age) VALUES(?,?)", "福建最美的妞", 66)
	tx.MustExec("UPDATE user SET age=16 WHERE name=?", "福建最美的妞")
	tx.Commit()
}

func main() {
	defer db.Close()
	//queryRow()
	//query()
	tx()
}
