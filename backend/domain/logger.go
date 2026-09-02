package domain

type Logger interface {
	Info(str string, args ...any)
	Error(str string, args ...any)
	Warn(str string, args ...any)
	Debug(str string, args ...any)
}
