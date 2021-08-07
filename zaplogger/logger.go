package zaplogger

import (
	"go.uber.org/zap"

	"git.corout.in/golibs/slog"
)

type loggerImpl struct {
	logger *zap.Logger
}

func New() slog.Logger {
	return &loggerImpl{}
}

func (l *loggerImpl) Error(args ...interface{}) {
	msg := l.getMessage(args...)
	l.logger.Error(msg.message, msg.fields...)
}

func (l *loggerImpl) Errorf(template string, args ...interface{}) {
	msg := l.getMessagef(template, args...)
	l.logger.Error(msg.message, msg.fields...)
}

func (l *loggerImpl) Fatal(args ...interface{}) {
	msg := l.getMessage(args...)
	l.logger.Fatal(msg.message, msg.fields...)
}

func (l *loggerImpl) Fatalf(template string, args ...interface{}) {
	msg := l.getMessagef(template, args...)
	l.logger.Fatal(msg.message, msg.fields...)
}

func (l *loggerImpl) Info(args ...interface{}) {
	msg := l.getMessage(args...)
	l.logger.Info(msg.message, msg.fields...)
}

func (l *loggerImpl) Infof(template string, args ...interface{}) {
	msg := l.getMessagef(template, args...)
	l.logger.Info(msg.message, msg.fields...)
}

func (l *loggerImpl) Panic(args ...interface{}) {
	msg := l.getMessage(args...)
	l.logger.Panic(msg.message, msg.fields...)
}

func (l *loggerImpl) Panicf(template string, args ...interface{}) {
	msg := l.getMessagef(template, args...)
	l.logger.Panic(msg.message, msg.fields...)
}

func (l *loggerImpl) Warn(args ...interface{}) {
	msg := l.getMessage(args...)
	l.logger.Warn(msg.message, msg.fields...)
}

func (l *loggerImpl) Warnf(template string, args ...interface{}) {
	msg := l.getMessagef(template, args...)
	l.logger.Warn(msg.message, msg.fields...)
}

func (l *loggerImpl) With(args ...interface{}) slog.Logger {
	msg := l.getMessage(args...)

	return &loggerImpl{
		logger: l.logger.With(msg.fields...),
	}
}

func (l *loggerImpl) Sync() {
	if err := l.logger.Sync(); err != nil {
		l.logger.Error("sync zap logger", zap.Error(err))
	}
}
