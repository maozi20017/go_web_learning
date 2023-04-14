package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPage 函數，用於顯示登入頁面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// RegisterPage 函數，用於顯示註冊頁面
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// IndexPage 函數，用於顯示內容頁面
func IndexPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	identity := session.Get("identity_description")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"username":             username,
		"identity_description": identity,
	})
}

// AdminPage 函數，用於顯示後台管理頁面
func AdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

func UserListPage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := GetUserList(db)
		c.HTML(http.StatusOK, "userlist.html", gin.H{
			"users": users,
		})
	}
}
