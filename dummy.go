package slog

func newDummy() Logger {
	return &dummyLogger{}
}

type dummyLogger struct{}

func (d dummyLogger) Error(...interface{}) {}

func (d dummyLogger) Errorf(string, ...interface{}) {}

func (d dummyLogger) Fatal(...interface{}) {}

func (d dummyLogger) Fatalf(string, ...interface{}) {}

func (d dummyLogger) Info(...interface{}) {}

func (d dummyLogger) Infof(string, ...interface{}) {}

func (d dummyLogger) Panic(...interface{}) {}

func (d dummyLogger) Panicf(string, ...interface{}) {}

func (d dummyLogger) Warn(...interface{}) {}

func (d dummyLogger) Warnf(string, ...interface{}) {}

func (d dummyLogger) With(...interface{}) Logger {
	return d
}

func (d dummyLogger) Sync() {}
