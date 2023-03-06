package main // 宣告這是一個 main package

import (
	"database/sql" // 引用 Go 語言的資料庫 SQL 操作函式庫
	"fmt"          // 引用 Go 語言的標準輸出函式庫

	_ "github.com/go-sql-driver/mysql" // 引用 MySQL 驅動程式
	// "_" 表示只要匯入該套件的 init() 函式，不使用該套件的其他函式
)

type User struct { // 宣告一個 User struct 結構體
	ID       string // 使用者 ID
	Username string // 使用者名稱
	Password string // 使用者密碼
}

func CreateTable(db *sql.DB) error { // 宣告 CreateTable 函式，建立資料庫 Table
	sql := `CREATE TABLE IF NOT EXISTS users(
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64)
	); ` // SQL 語句

	if _, err := db.Exec(sql); err != nil { // 執行 SQL 語句
		fmt.Println("建立 Table 發生錯誤:", err) // 輸出錯誤訊息
		return err
	}
	fmt.Println("建立 Table 成功！") // 輸出成功訊息
	return nil
}

func InsertUser(DB *sql.DB, username, password string) error { // 宣告 InsertUser 函式，插入一個使用者資料
	_, err := DB.Exec("insert INTO users(username,password) values(?,?)", username, password) // 執行 SQL 語句，插入一個使用者資料
	if err != nil {
		fmt.Printf("建立使用者失敗，原因是：%v", err) // 輸出錯誤訊息
		return err
	}
	fmt.Println("建立使用者成功！") // 輸出成功訊息
	return nil
}

func QueryUser(db *sql.DB, username string) { // 宣告 QueryUser 函式，查詢使用者資料
	user := new(User)                                                          // 創建一個 User struct 結構體
	row := db.QueryRow("select * from users where username=?", username)       // 執行 SQL 語句，查詢使用者資料
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil { // 讀取查詢結果，並映射到 User struct 結構體中
		fmt.Printf("映射使用者失敗，原因為：%v\n", err) // 輸出錯誤訊息
		return
	}
	fmt.Println("查詢使用者成功", *user) // 輸出成功訊息及使用者資料
}
