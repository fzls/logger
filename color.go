package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// 为了不引入新依赖，直接将对应库需要的部分复制过来了，具体可参考 github.com\gookit\color@v1.4.2\color.go

// ResetSet 重置色彩 ansi code
const ResetSet = "\x1b[0m"

const (
	// SettingTpl 开始色彩 ansi code
	SettingTpl = "\x1b[%sm"
)

var (
	colorCodePanic = fmt.Sprintf(SettingTpl, "1;31") // color.Style{color.Bold, color.Red}.String()
	colorCodeFatal = fmt.Sprintf(SettingTpl, "1;31") // color.Style{color.Bold, color.Red}.String()
	colorCodeError = fmt.Sprintf(SettingTpl, "31")   // color.Style{color.Red}.String()
	colorCodeWarn  = fmt.Sprintf(SettingTpl, "33")   // color.Style{color.Yellow}.String()
	colorCodeInfo  = fmt.Sprintf(SettingTpl, "32")   // color.Style{color.Green}.String()
	colorCodeDebug = fmt.Sprintf(SettingTpl, "37")   // color.Style{color.White}.String()
	colorCodeTrace = fmt.Sprintf(SettingTpl, "36")   // color.Style{color.Cyan}.String()
)

// GetLogLevelColorCode 获取日志等级对应色彩code
func GetLogLevelColorCode(level logrus.Level) string {
	switch level {
	case logrus.PanicLevel:
		return colorCodePanic
	case logrus.FatalLevel:
		return colorCodeFatal
	case logrus.ErrorLevel:
		return colorCodeError
	case logrus.WarnLevel:
		return colorCodeWarn
	case logrus.InfoLevel:
		return colorCodeInfo
	case logrus.DebugLevel:
		return colorCodeDebug
	case logrus.TraceLevel:
		return colorCodeTrace

	default:
		return colorCodeInfo
	}
}
