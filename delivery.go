package main

import (
	"errors"
	"go_web_learning/model"
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
	identity := session.Get("identity")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
		"identity": identity,
	})
}

// // AdminPage 函數，用於顯示後台管理頁面
// func AdminPage(c *gin.Context) {
// 	user := c.MustGet("user").(*model.User)
// 	if user.IdentityID != 3 {
// 		c.HTML(http.StatusUnauthorized, "unauthorized.html", nil)
// 		return
// 	}
// 	c.HTML(http.StatusOK, "admin.html", nil)
// }

// Login 函數，用於驗證用戶名和密碼是否正確
func Login(db *gorm.DB, username string, password string) error {
	user, err := FindUser(db, username)
	if err != nil {
		return errors.New("使用者不存在")
	}
	if user.Password != password {
		return errors.New("密碼錯誤")
	}

	return nil
}

// Register 函數，用於註冊新用戶，驗證用戶名和密碼是否正確
func Register(db *gorm.DB, username string, password string) error {
	_, err := FindUser(db, username)
	if err == nil {
		return errors.New("使用者已存在")
	} else {
		return CreateUser(db, &model.User{
			Username: username,
			Password: password,
		})
	}
}
