package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	viper.AddConfigPath("conf")   // 设置配置文件所在目录
	viper.SetConfigName("app") //设置配置文件名称
	viper.SetConfigType("yaml")   // 设置配置文件格式为YAML
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	runMode := viper.GetString("runmode")
	addr:=viper.GetString("addr")
	gin.SetMode(runMode) // 设置gin运行模式
	r := gin.Default()
	r.GET("test/config", func(c *gin.Context) {
		c.String(http.StatusOK,runMode)
	})
	r.Run(addr)
}
