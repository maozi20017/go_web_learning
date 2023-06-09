package main // 包名為 main，表明這是一個可執行的 Go 程序

import (
	"errors" // 引入 errors 包，用於返回錯誤
	"go_web_learning/model"
	"net/http" // 引入 net/http 包，用於處理 HTTP 請求

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin" // 引入 gin 框架
	"gorm.io/gorm"             // 引入 gorm 框架
)

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
			Username:   username,
			Password:   password,
			IdentityID: 1,
		})
	}
}

// ValidateForm 函數，用於驗證表單參數是否正確
func ValidateForm(c *gin.Context) (string, string, error) {
	var (
		username string
		password string
	)
	// 解析表單參數，檢查是否有 username 參數，並且值不為空
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		// 如果沒有 username 參數或值為空，返回一個帶有錯誤信息的 HTML 頁面
		return "", "", errors.New("必須輸入使用者名稱")
	}
	// 解析表單參數，檢查是否有 password 參數，並且值不為空
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		// 如果沒有 password 參數或值為空，返回一個帶有錯誤信息的 HTML 頁面
		return "", "", errors.New("必須輸入密碼")
	}
	return username, password, nil
}

// RegisterAuth 函數，用於處理註冊請求，驗證用戶名和密碼是否正確
func RegisterAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, err := ValidateForm(c)
		if err != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"error": err,
			})
			return
		}
		if err := Register(db, username, password); err == nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"success":  "註冊成功 2秒後將導向至登入頁面",
				"redirect": "/login",
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

// LoginAuth 函數，用於處理登入請求
func LoginAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, err := ValidateForm(c)
		if err != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{
				"error": err,
			})
			return
		}
		if err := Login(db, username, password); err == nil {
			// 設定 session 變數
			session := sessions.Default(c)
			session.Set("username", username)
			user := &model.User{}
			db.Where("username = ?", username).Preload("Identity").First(user)
			session.Set("identity_description", user.Identity.Description)
			session.Save()

			c.HTML(http.StatusOK, "login.html", gin.H{
				"success":  "登入成功 2秒後將導向至內容頁面",
				"redirect": "/index",
			})
			return
		} else {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"error": err,
			})
			return
		}
	}
}
