//使用go内置HTTP 服务器
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"demo01/router"

)

func main() {
	// 1、初始化引擎
	engine := gin.Default()
	// 2、注册一个路由和处理函数
	engine.Any("/", router.WebRoot)
	// 3、绑定端口，启动http服务
	server := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}



