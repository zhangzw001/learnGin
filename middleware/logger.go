package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	config "github.com/zhangzw001/learnGin/config/dev"
	"log"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.LogFilePath
	logFileName := config.LogFileName

	// 日志文件
	fileName := path.Join(logFilePath,logFileName)

	// 写入文件
	file , err := os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = file
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999999999Z07:00",
	})

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		during := endTime.Sub(startTime)
		// 请求方法
		reqMethod := c.Request.Method
		// 请求uri
		reqUri := c.Request.RequestURI
		//状态码
		code := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("[%3d %v %s] %s %s ",
			code,during,clientIP,reqMethod,reqUri)

		// 日志格式
		//logger.WithFields(logrus.Fields{
		//	"status_code"  : code,
		//	"latency_time" : during,
		//	"client_ip"    : clientIP,
		//	"req_method"   : reqMethod,
		//	"req_uri"      : reqUri,
		//}).Info()
	}
}
