/**
  @author: Zero
  @date: 2023/3/18 21:08:31
  @desc: 基于rotateLogs 实现日志文件的切割

**/

package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

// NewRotateLogs 日志文件切割规则
// logDir     日志存放目录   logs/xxx/
// fileName   日志文件名   user-service
// day	      间隔多久生成一个新的日志文件 单位为日
// maxSaveDay 日志文件最多保留多少日  30
func NewRotateLogs(logDir, fileName string, day, maxSaveDay int) (*rotatelogs.RotateLogs, error) {
	logFilePath := fmt.Sprintf("%s/%s", logDir, fileName)
	write, err := rotatelogs.New(
		logFilePath+".%Y-%m-%d.log",                                   //日志文件命名格式
		rotatelogs.WithLinkName(logFilePath+".log"),                   //生成正在写入的日志文件,方便查看
		rotatelogs.WithRotationTime(time.Duration(day*24)*time.Hour),  //日志切割时间间隔
		rotatelogs.WithMaxAge(time.Duration(maxSaveDay*24)*time.Hour), //日志最长保留时间
	)
	if err != nil {
		return nil, err
	}
	return write, nil
}
