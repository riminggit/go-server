package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-server-template/config"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

//初始化log
func InitLogging() *logrus.Logger {
	config := projectConfig.AppConfig

	// 日志文件
	fileName := path.Join(config.LogConfig.LOG_SAVE_URL, config.LogConfig.LOG_SAVE_NAME)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	Logger = logrus.New()

	Logger.SetReportCaller(true) // 显示行号等信息

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)
	//设置输出
	Logger.Out = src

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return Logger
}

func WriteMiddlewareLogs(c *gin.Context) {
	//开始时间
	startTime := time.Now()
	//处理请求
	c.Next()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//请求方式
	reqMethod := c.Request.Method
	//请求路由
	reqUrl := c.Request.RequestURI
	//状态码
	statusCode := c.Writer.Status()
	//请求ip
	clientIP := c.ClientIP()

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"status_code":  statusCode,
		"latency_time": latencyTime,
		"client_ip":    clientIP,
		"req_method":   reqMethod,
		"req_uri":      reqUrl,
	}).Info()
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Debug(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Info(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Warn(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Error(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Fatal(v)
}
