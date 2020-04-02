package main

import (
	"database/sql"
	"fmt"

	"github.com/gchaincl/dotsql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//获取数据库句柄
	db, _ := sql.Open("sqlite3", ":memory:")

	//加载SQL文件
	dot, _ := dotsql.LoadFromFile("data.sql")

	dot.Exec(db, "create-users-table")

	dot.Exec(db, "create-user", "User Name", "main@example.com")

	rows, _ := dot.Query(db, "find-users-by-email", "main@example.com")
	var (
		id    int
		name  string
		email string
	)
	for rows.Next() {
		rows.Scan(&id, &name, &email)
		fmt.Println(id, name, email)
	}
}
