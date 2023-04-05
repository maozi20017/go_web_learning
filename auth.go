package main // 包名為 main，表明這是一個可執行的 Go 程序

import (
	"errors"   // 引入 errors 包，用於返回錯誤
	"net/http" // 引入 net/http 包，用於處理 HTTP 請求

	"github.com/gin-gonic/gin" // 引入 gin 框架
	"gorm.io/gorm"
)

// Auth 函數，用於驗證用戶名和密碼是否正確
func Auth(db *gorm.DB, username string, password string) error {
	user, err := FindUser(db, username)
	if err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("password is not correct")
	}
	return nil
}

// LoginPage 函數，用於顯示登入頁面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginAuth 函數，用於處理登入請求
func LoginAuth(db *gorm.DB, c *gin.Context) {
	var (
		username string
		password string
	)
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	if err := Auth(db, username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}
