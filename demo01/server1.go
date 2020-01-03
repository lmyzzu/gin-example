//使用Gin默认服务器
package main

import (
	"github.com/gin-gonic/gin"

	"demo01/router"
)

func main() {
	// 1、初始化引擎
	engine := gin.Default()
	// 2、注册一个路由和处理函数
	engine.Any("/", router.WebRoot)
	// 3、绑定端口，启动http服务
	engine.Run(":8080")
}

