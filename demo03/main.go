//使用http.Server内置的Shutdown方法来实现实现优雅退出、平滑重启的服务器
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 1、初始化引擎（带有默认中间件）
	engine := gin.Default()
	//创建不带中间件的引擎
	//engine:= gin.New()

	// 2、注册路由
	// 1)路由参数
	engine.GET("/test/:number", func(context *gin.Context) {
		// 使用 context.Param(key) 获取 url 参数
		number := context.Param("number")
		context.String(http.StatusOK, "number=%s", number)
	})
	// 2)URL参数
	engine.GET("/test", func(context *gin.Context) {
		//若name字段为空，设置默认值为Guest
		name := context.DefaultQuery("name", "Guest")
		context.String(http.StatusOK, "name=%s", name)
	})
	// 3)表单参数
	engine.POST("/test", func(context *gin.Context) {
		name := context.DefaultPostForm("name", "Guest")
		age := context.PostForm("age")
		context.String(http.StatusOK, "name=%s age=%s", name, age)
	})
	// 4)上传文件
	engine.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// 5)json数据
	engine.POST("/add", func(context *gin.Context) {
		// 获取原始字节
		rawData, err := context.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		strData := string(rawData)
		context.String(http.StatusOK, "query data=%s", strData)
	})

	// 3、绑定端口，启动http服务
	server := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 4、阻塞等待中断信号以优雅地关闭服务器（设置5秒的超时时间）
	//信号接收通道
	signalChan := make(chan os.Signal)
	//将进入的信号转到signalChan
	signal.Notify(signalChan, os.Interrupt)
	//阻塞
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
