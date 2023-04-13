package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter 設定路由器並回傳 gin.Engine 物件
// db 為 GORM 資料庫連線物件
func SetupRouter(db *gorm.DB) *gin.Engine {
	// 建立 gin 伺服器物件
	server := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	// 載入 HTML 模板檔案
	server.LoadHTMLGlob("static/html/*")

	// 設定靜態檔案目錄
	server.Static("/static", "./static")

	// 設定路由器
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth(db))
	server.GET("/register", RegisterPage)
	server.POST("/register", RegisterAuth(db))
	server.GET("/index", IndexPage)
	// server.GET("/admin", AdminPage)

	// 回傳路由器
	return server

}
