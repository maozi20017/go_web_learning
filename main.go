package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper" // 引入 viper 套件
	"gorm.io/driver/mysql"   // 引入 gorm 支援的 mysql driver
	"gorm.io/gorm"           // 引入 gorm 套件
)

type (
	Config struct { // 定義 Config 結構
		Username string `yaml:"username"` // 使用 yaml 格式的註解，定義屬性名稱
		Password string `yaml:"password"`
		Network  string `yaml:"network"`
		Server   string `yaml:"server"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	}
)

func main() {

	// 設定 Viper 套件
	viper.SetConfigName("config")   // 設定配置檔案名稱
	viper.SetConfigType("yaml")     // 設定配置檔案格式
	viper.AddConfigPath("./config") // 設定配置檔案路徑

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

	if err := db.AutoMigrate(new(Role)); err != nil { // 自動建立 Role 資料表
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}

	// 建立一筆 Role_id 為 0 的 Role 紀錄，即一般使用者
	db.Create(&Role{Role_id: 1, Description: "一般使用者"})
	// 建立一筆 Role_id 為 1 的 Role 紀錄，即版主
	db.Create(&Role{Role_id: 2, Description: "版主"})
	// 建立一筆 Role_id 為 2 的 Role 紀錄，即管理員
	db.Create(&Role{Role_id: 3, Description: "管理員"})

	// 建立一個 Role_id 為 3 的版主使用者
	db.Create(&User{Username: "10811225", Password: "ghost8797", Role_ID: 3})

	// 創建一個 gin 引擎實例
	server := gin.Default()
	// 設置模板文件路徑
	server.LoadHTMLGlob("static/html/*")
	// 啟動服務，監聽 8888 端口
	//設定靜態資源的讀取
	server.Static("/static", "./static")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth(db))
	server.GET("/register", RegisterPage)
	server.POST("/register", RegisterAuth(db))
	server.GET("/index", IndexPage)
	server.Run(":8888")
}
