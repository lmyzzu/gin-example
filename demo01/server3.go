//使用http.Server内置的Shutdown方法来实现实现优雅退出、平滑重启的服务器
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}



