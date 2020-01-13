package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Static("/assets", "./assets")
	engine.Run(":8080")
}
