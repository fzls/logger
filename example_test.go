package logger

import (
	"fmt"
	"os"
	"testing"
)

// 2020/05/31 19:06 by fzls

func TestMain(m *testing.M) {
	logger, err := NewLogger("logs", "test", "debug")
	if err != nil {
		fmt.Printf("new logger err=%v\n", err)
		return
	}

	logger.Debugf("Debugf %v", 1)
	logger.Infof("Infof %v", 1)
	logger.Warnf("Warnf %v", 1)
	logger.Errorf("Errorf %v", 1)
	// logger.DPanicf("DPanicf %v", 1)
	// logger.Panicf("Panicf %v", 1)
	logger.Fatalf("Fatalf %v", 1)

	os.Exit(0)
}
