package util

import (
	"fmt"
	"testing"
)

func slog() {
	if true {
		LogD_(3, "slog %v", "Debug")
	}
}

func TestLog(t *testing.T) {
	Debug("%v", "Debug")
	Info("%v", "Info")
	Warning("%v", "Warning")
	NewError("%v", "NewError")

	logger.Debug("logger %v", "Debug")
	logger.Info("logger %v", "Info")
	logger.Warning("logger %v", "Warning")
	logger.Error("logger %v", "Error")

	LogD_(-2, "-2 %v", "Debug")
	LogD_(-1, "-1 %v", "Debug")
	LogD_(0, "0 %v", "Debug")
	LogD_(1, "1 %v", "Debug")
	LogD_(2, "2 %v", "Debug")

	slog()

	fmt.Println("测试日志完毕")
}
