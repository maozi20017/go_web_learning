package main // 宣告這是一個 main package

import (
	"gorm.io/gorm" // 引入 gorm 套件
)

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error // 在資料庫中新增一筆 user 資料
}

func FindUser(db *gorm.DB, username string) (*User, error) {
	user := new(User)        // 建立一個新的 User 變數
	user.Username = username // 將 User 變數的 ID 欄位設定為參數中指定的 id
	err := db.Where("username = ?", username).First(user).Error
	return user, err // 回傳查詢到的 user 資料及錯誤訊息
}
