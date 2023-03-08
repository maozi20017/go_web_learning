package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql" // 引入 gorm 支援的 mysql driver
	"gorm.io/gorm"         // 引入 gorm 套件
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Network  string `yaml:"network"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"` // 使用者名稱
	Password string `json:""`         // 密碼
}

func main() {
	// 設定 Viper 套件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 讀取 YAML 檔案
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("無法讀取配置檔案: %s", err))
	}

	// 解析配置檔案
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("無法解析配置檔案: %s", err))
	}

	// 轉換成數據庫連接字串
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Network, config.Server, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 透過 dsn 連接資料庫
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
	if err := db.AutoMigrate(new(User)); err != nil { // 自動建立 User 資料表
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}

	user := &User{
		Username: "test", // 使用者名稱
		Password: "test", // 密碼
	}
	if err := CreateUser(db, user); err != nil { // 在資料庫中新增一筆 user 資料
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}

	if user, err := FindUser(db, 1); err == nil { // 查詢符合條件的 user 資料
		log.Println("查詢到 User 為 ", user)
	} else {
		panic("查詢 user 失敗，原因為 " + err.Error())
	}
}
