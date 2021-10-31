package logger

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

// FileHook logrus本地钩子，写到文件中
type FileHook struct {
	lock      *sync.Mutex
	levels    []logrus.Level   // hook级别
	formatter logrus.Formatter // 格式
	path      string           // 写入path
}

// NewFileHook 初始化本地日志钩子实现
func NewFileHook(logFilePath string, fileFormatter logrus.Formatter, levels []logrus.Level) *FileHook {
	hook := &FileHook{
		lock: new(sync.Mutex),
	}
	hook.formatter = fileFormatter
	hook.levels = append(hook.levels, levels...)
	hook.path = logFilePath

	return hook
}

// Levels ref: logrus/hooks.go impl Hook interface
func (hook *FileHook) Levels() []logrus.Level {
	if len(hook.levels) == 0 {
		return logrus.AllLevels
	}
	return hook.levels
}

// Fire ref: logrus/hooks.go impl Hook interface
func (hook *FileHook) Fire(entry *logrus.Entry) error {
	hook.lock.Lock()
	defer hook.lock.Unlock()

	dir := filepath.Dir(hook.path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fd, err := os.OpenFile(hook.path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer fd.Close()

	log, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = fd.Write(log)
	return err
}
