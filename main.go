package main

import (
	"fmt"
	"log"
)

func main() {
	// 連接數據庫
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("無法連接數據庫: %s", err)
	}

	defer func() {
		db = nil
	}()

	// 查詢使用者資料，預先載入身分資訊
	var users []User
	if err := db.Preload("Identity").Where("identity_id = ?", 3).Find(&users).Error; err != nil {
		panic(err)
	}

	// 輸出使用者資訊
	for _, user := range users {
		fmt.Printf("User: %v, Identity: %v\n", user.Username, user.Identity.Description)
	}

	// 創建 gin 引擎實例
	server := SetupRouter(db)

	// 啟動服務，監聽 8888 端口
	server.Run(":8888")
}
