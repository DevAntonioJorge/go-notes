package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Logger struct {
	logger *log.Logger
	level  LogLevel
	mu     sync.Mutex
}

var (
	defaultLogger *Logger
	once          sync.Once
)

func New(out io.Writer, prefix string, flag int, level LogLevel) *Logger {
	return &Logger{
		logger: log.New(out, prefix, flag),
		level:  level,
	}
}

func Default() *Logger {
	once.Do(func() {
		defaultLogger = New(os.Stdout, "", log.LstdFlags, LevelInfo)
	})
	return defaultLogger
}

func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(LevelDebug, "DEBUG", format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(LevelInfo, "INFO ", format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(LevelWarn, "WARN ", format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(LevelError, "ERROR", format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(LevelFatal, "FATAL", format, v...)
	os.Exit(1)
}

func (l *Logger) log(level LogLevel, prefix, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	var msg string
	if len(v) > 0 {
		msg = fmt.Sprintf("["+prefix+"] "+format, v...)
	} else {
		msg = "[" + prefix + "] " + format
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if level == LevelFatal {
		l.logger.Fatal(msg)
	} else {
		l.logger.Println(msg)
	}
}

func SetLevel(level LogLevel) {
	Default().SetLevel(level)
}

func Debug(format string, v ...interface{}) {
	Default().Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	Default().Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	Default().Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	Default().Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	Default().Fatal(format, v...)
}
