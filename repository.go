package main // 宣告這是一個 main package

import (
	"go_web_learning/model"

	"gorm.io/gorm" // 引入 gorm 套件
)

func CreateUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error // 在資料庫中新增一筆 User 資料
}

func FindUser(db *gorm.DB, username string) (*model.User, error) {
	user := new(model.User)  // 建立一個新的 User 變數
	user.Username = username // 將 User 變數的 ID 欄位設定為參數中指定的 id
	err := db.Where("username = ?", username).First(user).Error
	return user, err // 回傳查詢到的 user 資料及錯誤訊息
}

func GetUserList(db *gorm.DB) ([]*model.User, error) {
	users := []*model.User{}
	err := db.Preload("Identity").Find(&users).Error
	return users, err
}

func DeleteUser(db *gorm.DB, user *model.User) error {
	return db.Delete(user).Error // 在資料庫中刪除指定的 user 資料
}
