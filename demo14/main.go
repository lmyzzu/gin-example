package main

import (
	. "demo14/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(Logger())
	r.GET("test/log", func(c *gin.Context) {
		Log.Info("hello")//只需要记录一条简单的消息，可以不使用Field机制
		//Field机制输出日志
		Log.WithFields(logrus.Fields{
			"name": "zsl10",
			"uid":  "323222423",
			"sex":  "male",
		}).WithFields(CommonFields).Info("log_user_info")
		c.String(http.StatusOK, "ok")
	})
	r.Run(":8080")
}
