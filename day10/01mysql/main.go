package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入并使用其 init() 函数
)

func main() {
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	dsn := "root:root@tcp(127.0.0.1:3306)/gofullstack"
	// DB 是一个数据库(操作)句柄，代表一个具有零到多个底层连接的连接池
	// 它可以安全的被多个 goroutine 同时使用
	db, err := sql.Open("mysql", dsn) // Open 不会校验用户名、密码是否正确
	if err != nil {                   // DSN 格式不对
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil { // 连接信息不对
		fmt.Println("Ping failed:", err)
	}
}
