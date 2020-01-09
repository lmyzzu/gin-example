package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	gin.DisableConsoleColor()	// 禁用控制台颜色
	f, _ := os.Create("gin_test.log")	// 日志写入文件
	gin.DefaultWriter = io.MultiWriter(f)
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)  //同时将日志写入文件和控制台
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义日志格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	r.GET("test/log", func(c *gin.Context) {
		tempMap := map[string]int{"a": 1, "b": 2, "c": 3}
		fmt.Fprintln(gin.DefaultWriter, "test log")
		fmt.Fprintln(gin.DefaultWriter,tempMap)
		c.String(http.StatusOK,"ok")
	})
	r.Run(":8080")
}
