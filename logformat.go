package logger

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// LogFormat specialize for go-cqhttp
type LogFormat struct {
	EnableColor bool
}

// Format implements logrus.Formatter
func (f LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	buf := NewBuffer()
	defer PutBuffer(buf)

	if f.EnableColor {
		buf.WriteString(GetLogLevelColorCode(entry.Level))
	}

	buf.WriteString(entry.Time.Format("2006-01-02 15:04:05.999"))
	buf.WriteString(" ")
	buf.WriteString(strings.ToUpper(entry.Level.String()))

	if entry.Caller != nil {
		buf.WriteString(" ")
		buf.WriteString(TrimmedPath(entry.Caller.File))
		buf.WriteString(":")
		buf.WriteString(strconv.FormatInt(int64(entry.Caller.Line), 10))
	}
	buf.WriteString(" ")
	buf.WriteString(entry.Message)
	buf.WriteString(" \n")

	if f.EnableColor {
		buf.WriteString(ResetSet)
	}

	ret := append([]byte(nil), buf.Bytes()...) // copy buffer
	return ret, nil
}

func TrimmedPath(filepath string) string {
	idx := strings.LastIndexByte(filepath, '/')
	if idx == -1 {
		return filepath
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(filepath[:idx], '/')
	if idx == -1 {
		return filepath
	}

	return filepath[idx+1:]
}
