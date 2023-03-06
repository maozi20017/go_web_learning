package main // 宣告 main package

import (
	"database/sql" // 引用 Go 標準庫 database/sql
	"fmt"          // 引用 Go 標準庫 fmt

	_ "github.com/go-sql-driver/mysql" // 引用 MySQL 驅動程式
)

const (
	USERNAME = "demo"      // 常數 USERNAME 為 demo
	PASSWORD = "demo123"   // 常數 PASSWORD 為 demo123
	NETWORK  = "tcp"       // 常數 NETWORK 為 tcp
	SERVER   = "127.0.0.2" // 常數 SERVER 為 127.0.0.2
	PORT     = 3306        // 常數 PORT 為 3306
	DATABASE = "demo"      // 常數 DATABASE 為 demo
)

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return
	}
	defer db.Close()

	InsertUser(db, "test", "test")
}
