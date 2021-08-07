package slog

type Logger interface {
	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})

	With(args ...interface{}) Logger

	Sync()
}