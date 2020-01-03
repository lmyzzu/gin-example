package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * 根请求处理函数
 * 所有本次请求相关的方法都封装在*gin.Context中
 */
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello world")
}
