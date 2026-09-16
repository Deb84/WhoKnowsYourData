package logger

import (
	"fmt"
	"time"
)

type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8

	LevelStringDebug = "DEBUG"
	LevelStringInfo  = "INFO"
	LevelStringWarn  = "WARN"
	LevelStringError = "ERROR"
)

type Args []any

type Config struct {
	Level Level
}

type Logger struct {
	config Config
}

func LevelString(level Level) string {
	switch level {
	case LevelDebug:
		return LevelStringDebug
	case LevelInfo:
		return LevelStringInfo
	case LevelWarn:
		return LevelStringWarn
	case LevelError:
		return LevelStringError
	default:
		return ""
	}
}

func LevelInt(level string) (Level, bool) {
	switch level {
	case LevelStringDebug:
		return LevelDebug, true
	case LevelStringInfo:
		return LevelInfo, true
	case LevelStringWarn:
		return LevelWarn, true
	case LevelStringError:
		return LevelError, true
	default:
		return 0, false
	}
}

func NewLogger(config *Config) *Logger {
	if config == nil {
		config = &Config{
			Level: LevelInfo,
		}
	}
	return &Logger{
		config: *config,
	}
}

func (logger *Logger) SetLevel(levelString string) error {
	level, ok := LevelInt(levelString)
	if !ok {
		return fmt.Errorf("incorrect level %q", levelString)
	}
	logger.config.Level = level
	return nil
}

func (log *Logger) format(str string, level Level, args Args) string {
	str = fmt.Sprintf(str, args...)

	now := time.Now().Format("2006-01-02 15:04:05.000")
	levelString := LevelString(level)
	return fmt.Sprintf("[%s] [%s] %s", now, levelString, str)
}

func (logger *Logger) print(str string, level Level, args []any) {
	if level < logger.config.Level {
		return
	}

	parsed := logger.format(str, level, args)

	fmt.Println(parsed)
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
