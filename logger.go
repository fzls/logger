package logger

import (
	"fmt"
	"path"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func InitLogger(dir string, name string, consoleLevel, fileLevel string) {
	logrus.SetReportCaller(true)

	// 设置写console
	logrus.SetLevel(GetLogLevel(consoleLevel))
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(LogFormat{EnableColor: true})

	// 设置写文件
	logFilePath := getLoggerPath(dir, name)
	fileFormatter := LogFormat{EnableColor: false}
	fileHook := NewFileHook(logFilePath, fileFormatter, GetLogLevels(consoleLevel))
	logrus.AddHook(fileHook)
}

func getLoggerPath(dir string, name string) string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	logFileName := fmt.Sprintf("%v_%4d_%02d_%02d_%02d_%02d_%02d.log", name, year, month, day, hour, min, sec)

	return path.Join(dir, logFileName)
}
