package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	// 1、字符串响应
	r.GET("/res/string", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	// 2、json响应
	// 1) 使用gin.H,本质是 map[string]interface{}
	r.GET("/res/json/1", func(c *gin.Context) {
		// 输出头格式为 application/json; charset=UTF-8 的 json 字符串
		c.JSON(http.StatusOK, gin.H{"msg": "json1", "status": 200})
	})
	// 2) 使用结构体响应
	r.GET("/res/json/2", func(c *gin.Context) {
		var user struct {
			Name   string
			Msg    string
			Number int
		}
		user.Name = "zsl10"
		user.Msg = "hello"
		user.Number = 123
		c.JSON(200, user)
	})
	// 3、XML响应
	r.GET("/res/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"user": "zsl10", "msg": "ok"})
	})
	// 4、YAML响应
	r.GET("/res/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zsl10", "msg": "ok"})
	})
	// 5、html模板渲染
	//加载模板
	r.LoadHTMLGlob("templates/*")
	//定义路由
	r.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello world",
		})
	})
	// 6、重定向
	r.GET("/tmp", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	// 7、同步异步
	// 1)异步
	r.GET("/async", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})
	// 2)同步
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	r.Run(":8080")

}
