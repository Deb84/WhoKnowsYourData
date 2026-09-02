package logger

import (
	"fmt"
	"time"
)

const (
	LevelDebug   = -8
	LevelVerbose = -4
	LevelInfo    = 0
	LevelWarn    = 4
	LevelError   = 8
	LevelFatal   = 12
)

type Args []any

type Logger struct{}

func LevelString(level int) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelVerbose:
		return "VERBOSE"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

func NewLogger() *Logger {

	return &Logger{}
}

func (log *Logger) format(str string, level int, args Args) string {
	str = fmt.Sprintf(str, args...)

	now := time.Now().Format("2006-01-02 15:04:05.000")
	levelString := LevelString(level)
	return fmt.Sprintf("[%s] [%s] %s", now, levelString, str)
}

func (log *Logger) print(str string, level int, args Args) {
	formatted := log.format(str, level, args)
	fmt.Println(formatted)
}

func (log *Logger) Info(str string, args ...any) {
	log.print(str, LevelInfo, args)
}

func (log *Logger) Error(str string, args ...any) {
	log.print(str, LevelError, args)
}

func (log *Logger) Warn(str string, args ...any) {
	log.print(str, LevelWarn, args)
}

func (log *Logger) Debug(str string, args ...any) {
	log.print(str, LevelDebug, args)
}
