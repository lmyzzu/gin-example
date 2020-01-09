package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.DisableConsoleColor()	// 禁用控制台颜色
	f, _ := os.Create("gin_test.log")	// 日志写入文件
	gin.DefaultWriter = io.MultiWriter(f)
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)  //同时将日志写入文件和控制台
	r := gin.Default()
	r.GET("test/log", func(c *gin.Context) {
		c.String(http.StatusOK,"ok")
	})
	r.Run(":8080")
}
