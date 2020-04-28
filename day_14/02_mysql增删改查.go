package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB, err := sql.Open("mysql", "root:")
	if err != nil {
		fmt.Println("sql.Open", err)
	}

	defer DB.Close()
	//err =DB.Ping()
	if err := DB.Ping(); err != nil {
		fmt.Println("DB.Ping()", err)
	}

	//sql :=`insert into exchange values (6,"ceshizhanghao@vip.qq","1231312","me")`  // TODO 插入数据（添加数据）
	//result, err:=DB.Exec(sql)  // 执行sql 语句
	//n,_:=result.RowsAffected()  //获取受影响的记录数
	//fmt.Println("受影响的记录数：",n)
	//if err !=nil{
	//	fmt.Println("DB.Exec",err)
	//}

	// TODO 执行预处理
	sql_data := [4][5]string{{"6", "ceshizhanghao@vip.qq", "1231312", "me"}, {"7", "ceshizhanghao1@vip.qq", "12313121", "me1"}}
	//fmt.Println(sql_data)
	stmt, err := DB.Prepare(`insert into exchange values (?,?)`)
	//db.Prepare("insert into stu values (?,?)")
	if err != nil {
		fmt.Println("DB.Prepare", err)
	}
	for _, s := range sql_data {
		stmt.Exec(s[0], s[1]) //调用预处理语句
		//fmt.Printf(s[0],s[1])
	}
	//fmt.Printf("%T\n",result)   //TODO sql.driverResult 类型
}
