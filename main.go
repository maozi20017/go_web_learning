package main

import (
	"fmt"
	"net/http"

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

	User struct { // 定義 User 結構
		ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
		Username string `json:"username"` // 使用 json 格式的註解，定義屬性名稱
		Password string `json:"password"` // 使用 json 格式的註解，定義屬性名稱
	}

	// IndexData 結構體，用於存儲首頁數據
	IndexData struct {
		Title   string // 標題
		Content string // 內容
	}
)

// test 函數，用於處理首頁請求
func test(c *gin.Context) {
	// 創建一個 IndexData 對象
	data := new(IndexData)
	data.Title = "首頁"        // 設置標題
	data.Content = "我的第一個首頁" // 設置內容
	// 返回 HTML 響應
	c.HTML(http.StatusOK, "index.html", data)
}

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

	// 創建一個 gin 引擎實例
	server := gin.Default()
	// 設置模板文件路徑
	server.LoadHTMLGlob("static/html/*")
	// 註冊一個處理 GET 請求的路由規則，當路由為 "/" 時，調用 test 函數處理請求
	server.GET("/", test)
	// 啟動服務，監聽 8888 端口
	//設定靜態資源的讀取
	server.Static("/static", "./static")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth(db))
	server.GET("/register", RegisterPage)
	server.POST("/register", RegisterAuth(db))
	server.Run(":8888")

	// user := &User{
	// 	Username: "test", // 使用者名稱
	// 	Password: "test", // 密碼
	// }
	// // 在資料庫中新增一筆 user 資料，如果發生錯誤就 panic
	// if err := CreateUser(db, user); err != nil {
	// 	panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	// }

	// // 查詢符合條件的 user 資料，如果發生錯誤就 panic
	// if user, err := FindUser(db, user.Username); err == nil {
	// 	log.Println("查詢到 User 為 ", user)
	// } else {
	// 	panic("查詢 user 失敗，原因為 " + err.Error())
	// }
}
