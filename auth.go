package main // 包名為 main，表明這是一個可執行的 Go 程序

import (
	"errors"   // 引入 errors 包，用於返回錯誤
	"net/http" // 引入 net/http 包，用於處理 HTTP 請求

	"github.com/gin-gonic/gin" // 引入 gin 框架
)

// UserData 變量，用於存儲用戶名和密碼
var UserData map[string]string

func init() {
	// 初始化 UserData 變量，添加一個用戶
	UserData = map[string]string{
		"test": "test",
	}
}

// CheckUserIsExist 函數，用於檢查用戶是否存在
func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username] // 判斷用戶名是否存在
	return isExist
}

// CheckPassword 函數，用於檢查密碼是否正確
func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil // 密碼正確，返回 nil
	} else {
		return errors.New("password is not correct") // 密碼不正確，返回錯誤
	}
}

// Auth 函數，用於驗證用戶名和密碼是否正確
func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password) // 驗證密碼是否正確
	} else {
		return errors.New("user is not exist") // 用戶不存在，返回錯誤
	}
}

// LoginPage 函數，用於顯示登入頁面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginAuth 函數，用於處理登入請求
func LoginAuth(c *gin.Context) {
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
	if err := Auth(username, password); err == nil {
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
