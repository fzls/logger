package logger

import "github.com/sirupsen/logrus"

// GetLogLevels 获取日志等级列表
//
// 可能的值有
//
// "trace","debug","info","warn","warn","error"
func GetLogLevels(level string) []logrus.Level {
	switch level {
	case "trace":
		return []logrus.Level{
			logrus.TraceLevel, logrus.DebugLevel,
			logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel,
			logrus.FatalLevel, logrus.PanicLevel,
		}
	case "debug":
		return []logrus.Level{
			logrus.DebugLevel, logrus.InfoLevel,
			logrus.WarnLevel, logrus.ErrorLevel,
			logrus.FatalLevel, logrus.PanicLevel,
		}
	case "info":
		return []logrus.Level{
			logrus.InfoLevel, logrus.WarnLevel,
			logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel,
		}
	case "warn":
		return []logrus.Level{
			logrus.WarnLevel, logrus.ErrorLevel,
			logrus.FatalLevel, logrus.PanicLevel,
		}
	case "error":
		return []logrus.Level{
			logrus.ErrorLevel, logrus.FatalLevel,
			logrus.PanicLevel,
		}
	default:
		return []logrus.Level{
			logrus.InfoLevel, logrus.WarnLevel,
			logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel,
		}
	}
}

// GetLogLevel 获取日志等级
//
// 可能的值有
//
// "trace","debug","info","warn","warn","error"
func GetLogLevel(level string) logrus.Level {
	switch level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
