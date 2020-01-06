package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
// binding:"required"：校验规则，不能为空
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 1) 绑定 url query
	router.GET("/test/url/query", func(c *gin.Context) {
		// 声明接收的变量
		var urlQuery Login
		// 将request的query中的数据，自动解析到结构体
		if err := c.ShouldBindQuery(&urlQuery); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, "Success")
	})

	// 2) 绑定表单参数
	router.POST("/form", func(c *gin.Context) {
		var form Login
		if err := c.Bind(&form); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, "Success")
	})

	// 3) 绑定json数据
	router.POST("/json", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, "Success")
	})




	router.Run(":8080")
}
