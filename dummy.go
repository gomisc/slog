package slog

func newDummy() Logger {
	return &dummyLogger{}
}

type dummyLogger struct{}

func (d dummyLogger) Error(...any) {}

func (d dummyLogger) Errorf(string, ...any) {}

func (d dummyLogger) Fatal(...any) {}

func (d dummyLogger) Fatalf(string, ...any) {}

func (d dummyLogger) Info(...any) {}

func (d dummyLogger) Infof(string, ...any) {}

func (d dummyLogger) Panic(...any) {}

func (d dummyLogger) Panicf(string, ...any) {}

func (d dummyLogger) Warn(...any) {}

func (d dummyLogger) Warnf(string, ...any) {}

func (d dummyLogger) Debug(...any) {}

func (d dummyLogger) Debugf(string, ...any) {}

func (d dummyLogger) With(...any) Logger {
	return d
}

func (d dummyLogger) Sync() {}
