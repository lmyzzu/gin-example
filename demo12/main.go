package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/spf13/pflag"
	"log"
	"net/http"
)

var env = pflag.StringP("env", "e", "", "set environment")

func main() {
	pflag.Parse()
	var env = *env

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	if env == "" {
		a := cfg.Section("env")
		b, _ := a.GetKey("value")
		env = b.String()
	}
	section := cfg.Section(env)
	addr, err := section.GetKey("addr")
	runMode, err := section.GetKey("runMode") //dev、test环境可以从父分区获取
	r := gin.Default()
	r.GET("test/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"addr": addr.String(),
			"mode": runMode.String(),
		})
	})
	r.Run(addr.String())
}
