package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

// 包含绑定数据和校验规则
type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	//gtfield=CheckIn:大于check_in字段的值
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//自定义验证规则
func bookableDate(fl validator.FieldLevel) bool {
	//获取参数值并转换为时间格式
	date := fl.Field().Interface().(time.Time)
	today := time.Now()
	if today.Unix() > date.Unix() {
		return false
	}
	return true
}

func main() {
	route := gin.Default()
	//注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//绑定第一个参数是校验规则的函数第二个参数是自定义的验证函数
		v.RegisterValidation("bookabledate", bookableDate)
	}
	route.GET("/date", getBookable)
	route.Run()
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindQuery(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
