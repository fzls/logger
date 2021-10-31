package logger

import (
	"os"
	"testing"

	logger "github.com/sirupsen/logrus"
)

// 2020/05/31 19:06 by fzls

func TestMain(m *testing.M) {
	InitLogger("logs", "test", "info", "debug")

	logger.Debugf("Debugf %v", 1)
	logger.Infof("Infof %v", 1)
	logger.Warnf("Warnf %v", 1)
	logger.Errorf("Errorf %v", 1)
	// logger.DPanicf("DPanicf %v", 1)
	// logger.Panicf("Panicf %v", 1)
	logger.Fatalf("Fatalf %v", 1)

	os.Exit(0)
}
