package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // 是一个连接池，并发安全

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gofullstack"
	var err error
	// 非常重要：不要使用 := 来创建局部变量
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open db failed:", err)
		return
	}
	db.SetMaxOpenConns(1) // 设置连接池最大连接数
	db.SetMaxIdleConns(1)
}

type user struct {
	id   uint32
	name string
	age  uint8
}

// queryRow 查询单行
func queryRow(id uint32) {
	sqlStr := "SELECT id, name, age FROM user WHERE id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
	if err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age); err != nil {
		fmt.Println("QueryRow failed:", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
}

// query 多行查询
func query() {
	sqlStr := "SELECT id,name,age FROM user ORDER BY id DESC"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}
	// 非常重要
	defer rows.Close()
	var us []user
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.id, &u.name, &u.age); err != nil {
			fmt.Println("Scan failed:", err)
			continue
		}
		us = append(us, u)

	}
	for _, u := range us {
		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}
}

// insert 插入记录
func insert(u *user) {
	sqlStr := "INSERT INTO user (name, age) VALUES(?,?)"
	result, err := db.Exec(sqlStr, u.name, u.age)
	if err != nil {
		fmt.Println("insert failed:", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get LastInsertId failed:", err)
		return
	}
	u.id = uint32(id)
	fmt.Println("insert success:", u)
}

// update 更新记录
func update(u *user) {
	sqlStr := "UPDATE user SET name=?, age=? WHERE id=?"
	result, err := db.Exec(sqlStr, u.name, u.age, u.id)
	if err != nil {
		fmt.Println("update failed:", err)
		return
	}
	aff, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get RowsAffected failed:", err)
		return
	}
	fmt.Println("更新的记录数：", aff)
}

func del(id uint32) {
	sqlStr := "DELETE FROM user WHERE id=?"
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed:", err)
		return
	}
	aff, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get RowsAffected failed:", err)
		return
	}
	fmt.Println("删除成功：", aff)
}
func main() {
	defer db.Close()
	//queryRow(1)
	//insert(&user{name:"隔壁王叔",age:38})
	//update(&user{id: 5, name: "王叔叔在隔壁", age: 40})
    del(5)
	query()
}
