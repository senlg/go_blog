package core

import (
	"bytes"
	"fmt"
	"go_blog/global"
	"io"
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 41
	yellow = 43
	blue   = 46
	green  = 42
)

type LogFormatter struct{}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = green
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue

	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006/01/02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] [%v]\x1b[%d;37m%-7s\x1b[0m %s %s %s \n", global.Config.Logger.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] [%v]\x1b[%d;37m%-7s\x1b[0m %s \n", global.Config.Logger.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
func InitLogger() *logrus.Logger {
	mlog := logrus.New()
	flog, err := os.OpenFile("logs/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Panic(err.Error())
	}
	logOutArr := io.MultiWriter([]io.Writer{os.Stdout, flog}...)
	mlog.SetOutput(logOutArr) // 设置输出类型
	mlog.SetReportCaller(global.Config.Logger.ShowLine)
	mlog.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level)
	return mlog
}
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout) // 设置输出类型
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
