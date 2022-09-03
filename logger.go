package slog

// Logger - интерфейс логгера
type Logger interface {
	Error(args ...any)
	Errorf(template string, args ...any)

	Fatal(args ...any)
	Fatalf(template string, args ...any)

	Info(args ...any)
	Infof(template string, args ...any)

	Panic(args ...any)
	Panicf(template string, args ...any)

	Warn(args ...any)
	Warnf(template string, args ...any)

	Debug(args ...any)
	Debugf(template string, args ...any)

	With(args ...any) Logger

	Sync()
}
