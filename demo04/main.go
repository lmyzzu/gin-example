package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Blog struct {
	ID int `uri:"tag_id" binding:"required"`
	Name string `uri:"blog_name" binding:"required"`
}

func main() {
	router := gin.Default()

	// 1) 绑定 url query
	router.GET("/test/url/query", func(c *gin.Context) {
		var urlQuery Login
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
		if err := c.BindJSON(&json); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, "Success")
	})

	// 4)绑定



	router.Run(":8080")
}
