package logger

// 2020/05/31 19:00 by fzls
import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel  = "debug"
	InfoLevel   = "info"
	WarnLevel   = "warn"
	ErrorLevel  = "error"
	DPanicLevel = "dpanic"
	PanicLevel  = "panic"
	FatalLevel  = "fatal"
)

var levelStrToZapLevel = map[string]zapcore.Level{
	DebugLevel:  zapcore.DebugLevel,
	InfoLevel:   zapcore.InfoLevel,
	WarnLevel:   zapcore.WarnLevel,
	ErrorLevel:  zapcore.ErrorLevel,
	DPanicLevel: zapcore.DPanicLevel,
	PanicLevel:  zapcore.PanicLevel,
	FatalLevel:  zapcore.FatalLevel,
}

func NewLogger(dir string, name string, level string) (*zap.SugaredLogger, error) {
	logLevel, ok := levelStrToZapLevel[strings.ToLower(level)]
	if !ok {
		return nil, fmt.Errorf("level invalid")
	}

	// 创建日志文件
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	logFileName := fmt.Sprintf("%v_%4d_%02d_%02d_%02d_%02d_%02d.log", name, year, month, day, hour, min, sec)
	logFilePath := path.Join(dir, logFileName)

	var logFile *os.File
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		logFile, err = os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_EXCL|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
	}

	// 初始化日志系统
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		CallerKey:   "linenum",
		MessageKey:  "msg",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder, // 小写编码器
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05.999"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(logLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                                          // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 构造日志
	logger := zap.New(core, caller, development).Sugar()

	return logger, nil
}
