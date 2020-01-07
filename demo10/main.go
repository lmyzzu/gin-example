package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("login", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "2131"
			//  maxAge int, 单位秒，有效期
			// path,cookie所在目录
			// domain string,域名
			// secure 是否能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})
	r.Run(":8080")
}
