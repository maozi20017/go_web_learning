package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("無法連接數據庫: %s", err)
	}

	defer func() {
		db = nil
	}()

	var user User
	db.Preload("Role").First(&user, 3)
	fmt.Printf("%+v", user)

	// 創建 gin 引擎實例
	server := SetupRouter(db)

	// 啟動服務，監聽 8888 端口
	server.Run(":8888")
}
