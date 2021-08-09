package zaplogger

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"git.corout.in/golibs/slog"
)

const (
	msgKey   = "reason"
	levelKey = "level"
	nameKey  = "logger"
)

type loggerImpl struct {
	logger *zap.Logger
}

// New - returns new slog.Logger like zap shugarred logger wrapper
func New(level int, colorize bool) slog.Logger {
	var (
		writer io.Writer = os.Stdout
		lvlEncoder = zapcore.CapitalLevelEncoder
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
				EncodeLevel:    lvlEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			}),
			zapcore.AddSync(writer),
			zap.NewAtomicLevelAt(zapLevel(level)),
		)),
	}
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

func zapLevel(lvl int) zapcore.Level {
	switch lvl {
	case slog.DebugLevel:
		return zapcore.DebugLevel
	case slog.InfoLevel:
		return zapcore.InfoLevel
	case slog.WarningLevel:
		return zapcore.WarnLevel
	case slog.ErrorLevel:
		return zapcore.ErrorLevel
	case slog.PanicLevel:
		return zapcore.PanicLevel
	case slog.FatalLevel:
		return zapcore.FatalLevel
	}

	return zapcore.FatalLevel
}
