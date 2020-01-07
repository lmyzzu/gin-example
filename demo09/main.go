package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过c.Get()获取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	{
		r.GET("/md3",MiddleWare(),func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{"request": req})
		})

		r.GET("/test",func(c *gin.Context) {
			c.String(http.StatusOK,"ok2")
		})
	}
	r.Run(":8080")
}