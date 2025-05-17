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

// New creates a new logger instance
func New(out io.Writer, prefix string, flag int, level LogLevel) *Logger {
	return &Logger{
		logger: log.New(out, prefix, flag),
		level:  level,
	}
}

// Default returns a default logger instance (stdout, no prefix, with timestamp, Info level)
func Default() *Logger {
	once.Do(func() {
		defaultLogger = New(os.Stdout, "", log.LstdFlags, LevelInfo)
	})
	return defaultLogger
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Debug logs a message at debug level
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(LevelDebug, "DEBUG", format, v...)
}

// Info logs a message at info level
func (l *Logger) Info(format string, v ...interface{}) {
	l.log(LevelInfo, "INFO ", format, v...)
}

// Warn logs a message at warning level
func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(LevelWarn, "WARN ", format, v...)
}

// Error logs a message at error level
func (l *Logger) Error(format string, v ...interface{}) {
	l.log(LevelError, "ERROR", format, v...)
}

// Fatal logs a message at fatal level and exits the program
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(LevelFatal, "FATAL", format, v...)
	os.Exit(1)
}

// log is the internal logging function
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

// Package level functions that use the default logger

// SetLevel sets the log level for the default logger
func SetLevel(level LogLevel) {
	Default().SetLevel(level)
}

// Debug logs a message at debug level using the default logger
func Debug(format string, v ...interface{}) {
	Default().Debug(format, v...)
}

// Info logs a message at info level using the default logger
func Info(format string, v ...interface{}) {
	Default().Info(format, v...)
}

// Warn logs a message at warning level using the default logger
func Warn(format string, v ...interface{}) {
	Default().Warn(format, v...)
}

// Error logs a message at error level using the default logger
func Error(format string, v ...interface{}) {
	Default().Error(format, v...)
}

// Fatal logs a message at fatal level and exits the program using the default logger
func Fatal(format string, v ...interface{}) {
	Default().Fatal(format, v...)
}
