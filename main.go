package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexData 結構體，用於存儲首頁數據
type IndexData struct {
	Title   string // 標題
	Content string // 內容
}

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
	// 創建一個 gin 引擎實例
	server := gin.Default()
	// 設置模板文件路徑
	server.LoadHTMLGlob("template/*")
	// 註冊一個處理 GET 請求的路由規則，當路由為 "/" 時，調用 test 函數處理請求
	server.GET("/", test)
	// 啟動服務，監聽 8888 端口
	server.Run(":8888")
}
