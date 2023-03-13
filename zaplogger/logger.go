package zaplogger

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/gomisc/slog.v1"
)

const (
	msgKey   = "reason"
	levelKey = "level"
	nameKey  = "logger"
	timeKey  = "ts"
)

type loggerImpl struct {
	logger *zap.Logger
}

// New - returns new slog.Logger like zap shugarred logger wrapper
// deprecated: use NewDevelopment or NewProduction instead
func New(level int, colorize bool) slog.Logger {
	var (
		writer     io.Writer = os.Stdout
		lvlEncoder           = zapcore.CapitalLevelEncoder
	)

	if colorize {
		lvlEncoder = zapcore.CapitalColorLevelEncoder
		writer = colorable.NewColorableStdout()
	}

	return &loggerImpl{
		logger: zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
				MessageKey:     msgKey,
				LevelKey:       levelKey,
				NameKey:        nameKey,
				TimeKey:        timeKey,
				EncodeLevel:    lvlEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			}),
			zapcore.AddSync(writer),
			zap.NewAtomicLevelAt(CountdownLevel(level)),
		)),
	}
}

func NewConsole(level int, colorize bool) slog.Logger {
	var (
		writer     io.Writer = os.Stdout
		lvlEncoder           = zapcore.CapitalLevelEncoder
	)

	if colorize {
		lvlEncoder = zapcore.CapitalColorLevelEncoder
		writer = colorable.NewColorableStdout()
	}

	return &loggerImpl{
		logger: zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
				MessageKey:     msgKey,
				LevelKey:       levelKey,
				NameKey:        nameKey,
				TimeKey:        timeKey,
				EncodeLevel:    lvlEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			}),
			zapcore.AddSync(writer),
			zap.NewAtomicLevelAt(CountdownLevel(level)),
		)),
	}
}

func NewJSON(level int) slog.Logger {
	return &loggerImpl{
		logger: zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey:     msgKey,
				LevelKey:       levelKey,
				NameKey:        nameKey,
				TimeKey:        timeKey,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			}),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(CountdownLevel(level)),
		)),
	}
}

func ZapFromSLog(log slog.Logger) *zap.Logger {
	if zl, ok := log.(*loggerImpl); ok {
		return zl.logger
	}

	return zap.NewNop()
}

func (l *loggerImpl) Error(args ...any) {
	msg := getMessage(args...)
	l.logger.Error(msg.message, msg.fields...)
}

func (l *loggerImpl) Errorf(template string, args ...any) {
	msg := getMessagef(template, args...)
	l.logger.Error(msg.message, msg.fields...)
}

func (l *loggerImpl) Fatal(args ...any) {
	msg := getMessage(args...)
	l.logger.Fatal(msg.message, msg.fields...)
}

func (l *loggerImpl) Fatalf(template string, args ...any) {
	msg := getMessagef(template, args...)
	l.logger.Fatal(msg.message, msg.fields...)
}

func (l *loggerImpl) Info(args ...any) {
	msg := getMessage(args...)
	l.logger.Info(msg.message, msg.fields...)
}

func (l *loggerImpl) Infof(template string, args ...any) {
	msg := getMessagef(template, args...)
	l.logger.Info(msg.message, msg.fields...)
}

func (l *loggerImpl) Panic(args ...any) {
	msg := getMessage(args...)
	l.logger.Panic(msg.message, msg.fields...)
}

func (l *loggerImpl) Panicf(template string, args ...any) {
	msg := getMessagef(template, args...)
	l.logger.Panic(msg.message, msg.fields...)
}

func (l *loggerImpl) Warn(args ...any) {
	msg := getMessage(args...)
	l.logger.Warn(msg.message, msg.fields...)
}

func (l *loggerImpl) Warnf(template string, args ...any) {
	msg := getMessagef(template, args...)
	l.logger.Warn(msg.message, msg.fields...)
}

func (l *loggerImpl) Debug(args ...any) {
	msg := getMessage(args...)
	l.logger.Debug(msg.message, msg.fields...)
}

func (l *loggerImpl) Debugf(template string, args ...any) {
	msg := getMessagef(template, args...)
	l.logger.Debug(msg.message, msg.fields...)
}

func (l *loggerImpl) With(args ...any) slog.Logger {
	msg := getMessage(args...)

	return &loggerImpl{
		logger: l.logger.With(msg.fields...),
	}
}

func (l *loggerImpl) Sync() {
	if err := l.logger.Sync(); err != nil {
		l.logger.Error("sync zap logger", zap.Error(err))
	}
}
