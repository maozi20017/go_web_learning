package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	server := gin.Default()
	server.LoadHTMLGlob("static/html/*")
	server.Static("/static", "./static")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth(db))
	server.GET("/register", RegisterPage)
	server.POST("/register", RegisterAuth(db))
	server.GET("/index", IndexPage)
	return server
}
