/**
  @author: Zero
  @date: 2023/3/18 21:00:21
  @desc: 日志组件初始化

**/

package initialize

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"zero.com/gtask-common/log"
	"zero.com/gtask-common/session"
)

// InitLogger 初始化Logrus和全局Logrus
func InitLogger() {
	session.Logger = initLogrus()
}

// initLogrus 初始化Logrus实例
func initLogrus() *logrus.Logger {
	//创建logrus一个实例
	logger := logrus.New()
	// 创建日志切割配置
	logFileWrite, err := log.NewRotateLogs("logs/", "user-service", 1, 30)
	if err != nil {
		panic(err)
	}
	//开启日志返回函数和行号
	logger.SetReportCaller(true)
	//TODO 因为终端日志部分内容 携带了color,所以要和日志文件采用不同的日志格式

	// 设置自定义的终端日志格式
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.ConsoleLogFormatter{})
	// 设置要输出的日志文件，和日志文件输出格式
	logger.AddHook(log.NewWriterHook(logFileWrite, &log.FileLogFormatter{}))
	// 设置logrus的日志级别
	logger.SetLevel(logrus.DebugLevel) //日志级别
	initGlobalLogrus(logFileWrite)
	return logger
}

// initGlobalLogrus 初始化全局Logrus
func initGlobalLogrus(logFileWrite *rotatelogs.RotateLogs) {
	// 是否开启日志代码行号
	logrus.SetReportCaller(true)
	// 设置全年局logrus终端日志格式
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&log.ConsoleLogFormatter{})
	// 设置全局logrus日志文件输出格式
	logrus.AddHook(log.NewWriterHook(logFileWrite, &log.FileLogFormatter{}))
	// 设置全局logrus日志级别
	logrus.SetLevel(logrus.DebugLevel)
}
