package log

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
	D("%v", "Debug")
	I("%v", "Info")
	W("%v", "Warning")
	E("%v", "NewError")

	logger.D("logger %v", "Debug")
	logger.I("logger %v", "Info")
	logger.W("logger %v", "Warning")
	logger.E("logger %v", "Error")

	LogD_(-2, "-2 %v", "Debug")
	LogD_(-1, "-1 %v", "Debug")
	LogD_(0, "0 %v", "Debug")
	LogD_(1, "1 %v", "Debug")
	LogD_(2, "2 %v", "Debug")

	slog()

	fmt.Println("测试日志完毕")
}
