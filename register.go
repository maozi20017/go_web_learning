package main // 包名為 main，表明這是一個可執行的 Go 程序

import (
	"errors"   // 引入 errors 包，用於返回錯誤
	"net/http" // 引入 net/http 包，用於處理 HTTP 請求

	"github.com/gin-gonic/gin" // 引入 gin 框架
	"gorm.io/gorm"             // 引入 gorm 框架
)

// Auth 函數，用於驗證用戶名和密碼是否正確
func Register(db *gorm.DB, username string, password string) error {
	_, err := FindUser(db, username)
	if err == nil {
		return errors.New("使用者已存在")
	} else {
		return CreateUser(db, &User{
			Username: username,
			Password: password,
		})
	}
}

// RegisterPage 函數，用於顯示登入頁面
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// LoginAuth 函數，用於處理登入請求
func RegisterAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			username string
			password string
		)
		// 解析表單參數，檢查是否有 username 參數，並且值不為空
		if in, isExist := c.GetPostForm("username"); isExist && in != "" {
			username = in
		} else {
			// 如果沒有 username 參數或值為空，返回一個帶有錯誤信息的 HTML 頁面
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"error": errors.New("必須輸入使用者名稱"),
			})
			return
		}
		// 解析表單參數，檢查是否有 password 參數，並且值不為空
		if in, isExist := c.GetPostForm("password"); isExist && in != "" {
			password = in
		} else {
			// 如果沒有 password 參數或值為空，返回一個帶有錯誤信息的 HTML 頁面
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"error": errors.New("必須輸入密碼"),
			})
			return
		}
		// 驗證用戶名和密碼是否正確，如果正確，返回一個帶有成功信息的 HTML 頁面，否則返回一個帶有錯誤信息的 HTML 頁面
		if err := Register(db, username, password); err == nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"success": "註冊成功",
			})
			return
		} else {
			c.HTML(http.StatusUnauthorized, "register.html", gin.H{
				"error": err,
			})
			return
		}
	}
}
