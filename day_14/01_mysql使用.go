package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//TODO 用户名 密码@[连接方式]（主机，端口号）/数据库名
	db, err := sql.Open("mysql", "")
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	err = db.Ping() //连接数据库
	if err != nil {
		fmt.Println("db.Ping", err)
		return
	}
	defer db.Close() // Todo 关闭数据库

}
