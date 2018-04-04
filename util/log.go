package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type logLevel int

type Logger struct {
	Level logLevel
	Log   *log.Logger
}

const (
	DEBUG   logLevel = iota
	INFO             = 2
	WARNING          = 3
	ERROR            = 4
)

// 包内私有对象
var logger = NewLogger()

// NewLogger返回一个新log对象
func NewLogger() *Logger {
	return &Logger{
		Log: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// SetLevel设置日子的打印等级
func (l *Logger) SetLevel(level int) {
	l.Level = logLevel(level)
}

// Redirect日志重定向,可以指定到控制台,文件等
func Redirect(writer io.Writer) {
	logger.Log.SetOutput(writer)
}

// RedirectFile日志重定向到文件,可以创建多层文件路径,如 /var/log/server/log/log.log
func RedirectFile(file string) error {
	var err error
	fp := filepath.Dir(file)
	err = os.MkdirAll(fp, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0755)
	if err != nil {
		return err
	}
	Redirect(f)
	return nil
}

// Debug message
func Debug(format string, args ...interface{}) {
	if DEBUG < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[D] "+format, args...))
}

// Info message
func Info(format string, args ...interface{}) {
	if INFO < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[I] "+format, args...))
}

// Warning message
func Warning(format string, args ...interface{}) {
	if WARNING < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[W] "+format, args...))
}

// Error message
func Error(format string, args ...interface{}) {
	if ERROR < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[E] "+format, args...))
}

// Debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if DEBUG < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[D] "+format, args...))
}

// Info message
func (l *Logger) Info(format string, args ...interface{}) {
	if INFO < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[I] "+format, args...))
}

// Warning message
func (l *Logger) Warning(format string, args ...interface{}) {
	if WARNING < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[W] "+format, args...))
}

// Error message
func (l *Logger) Error(format string, args ...interface{}) {
	if ERROR < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[E] "+format, args...))
}

func LogD_(callpath int, format string, args ...interface{}) {
	logger.Log.Output(callpath, fmt.Sprintf("[D] "+format, args...))
}
