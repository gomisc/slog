package slog

func newDummy() Logger {
	return &dummyLogger{}
}

type dummyLogger struct {}

func (d dummyLogger) Error(args ...interface{}) {}

func (d dummyLogger) Errorf(template string, args ...interface{}) {}

func (d dummyLogger) Fatal(args ...interface{}) {}

func (d dummyLogger) Fatalf(template string, args ...interface{}) {}

func (d dummyLogger) Info(args ...interface{}) {}

func (d dummyLogger) Infof(template string, args ...interface{}) {}

func (d dummyLogger) Panic(args ...interface{}) {}

func (d dummyLogger) Panicf(template string, args ...interface{}) {}

func (d dummyLogger) Warn(args ...interface{}) {}

func (d dummyLogger) Warnf(template string, args ...interface{}) {}

func (d dummyLogger) With(args ...interface{}) Logger {
	return d
}

func (d dummyLogger) Sync() {}


