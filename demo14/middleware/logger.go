package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"os"
	"time"

	//rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	//"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	//"time"
)

var Log *logrus.Logger
var CommonFields logrus.Fields
func Logger() gin.HandlerFunc {

	apiLogPath := "api.log"
	//禁止logrus的输出
	Log = logrus.New()
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	Log.Out = src
	//  只输出不低于当前级别是日志数据
	Log.SetLevel(logrus.DebugLevel)
	// 设置日志格式为json格式
	formatter := &logrus.JSONFormatter{
		PrettyPrint:     false,                 //美化json输出
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	}
	Log.SetFormatter(formatter) //设置日志格式
	Log.SetReportCaller(true)   //记录日志行号
	//输出文件
	logfile, _ := os.OpenFile(apiLogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	Log.SetOutput(logfile) //默认为os.stderr
	//日志分割
	logWriter, _ := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath), // 生成软链，指向最新日志文件
		//rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, formatter)
	logrus.AddHook(lfHook)//添加相应的hook
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start).String()

		//path := c.Request.URL.Path

		clientIP := c.ClientIP()
		//method := c.Request.Method
		statusCode := c.Writer.Status()

		//记录项目中要记录的通用字段
		CommonFields=logrus.Fields{
			"user_ip": clientIP,
			"latency": latency,
			"status_code":statusCode,
		}
		//
		//Log.Info(statusCode, "|",
		//	latency, "|",
		//	clientIP, "|",
		//	method, "|", path,
		//)
	}

}
