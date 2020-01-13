package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func calHandler(c *gin.Context) {
	var resContainer, sum int
	var success, resChan = make(chan int), make(chan int, 3)
	//设置超时
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	// 业务逻辑
	go MyLogic(resChan, success)
	for {
		select {
		case resContainer = <-resChan:
			sum += resContainer
			fmt.Println("add", resContainer)
		case <- success:
			c.JSON(http.StatusOK, gin.H{"code":200, "result": sum})
			return
		case <- ctx.Done(): //获取到一个超时的channel通知
			c.JSON(http.StatusOK, gin.H{"code":200, "result": sum})
			return
		}
	}
}

func main() {
	r := gin.New()
	r.GET("/calculate", calHandler)
	r.Run(":8080")
}

func MyLogic(rc chan<- int, success chan<- int) {
	wg := sync.WaitGroup{} // 创建一个 waitGroup 组
	wg.Add(3) // 运行3个协程
	go func() {
		rc <- a()
		wg.Done() // 完成一个，Done()一个
	}()

	go func() {
		rc <- b()
		wg.Done()
	}()

	go func() {
		rc <- c()
		wg.Done()
	}()

	wg.Wait() // 阻塞代码的运行，直到计数器的值减为0(等待3个协程执行完毕)
	success <- 1 // 发送一个成功信号到通道中
}

func a() int {
	time.Sleep(1*time.Second)
	return 1
}

func b() int {
	time.Sleep(2*time.Second)
	return 2
}

func c() int {
	time.Sleep(6*time.Second)
	return 3
}
