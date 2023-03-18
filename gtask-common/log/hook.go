/**
  @author: Zero
  @date: 2023/3/18 21:05:49
  @desc:

**/

package log

import (
	"github.com/sirupsen/logrus"
	"io"
)

// WriterHook 自定义Hook
// 用于将日志输入到不同Writer中时使用不同的日志格式
type WriterHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
}

// NewWriterHook 构建WriterHook
// writer: 要写入的日志文件
// formatter 日志文件的格式
func NewWriterHook(writer io.Writer, formatter logrus.Formatter) *WriterHook {
	return &WriterHook{
		Writer:    writer,
		Formatter: formatter,
	}
}

// Levels 设置哪些级别的日志会触发此钩子函数
func (hook *WriterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 日志数据写入到指定writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}
