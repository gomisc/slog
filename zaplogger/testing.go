package zaplogger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"git.corout.in/golibs/slog"
)

// Testing - logger соответсвующий logger.Logger и zaptest.TestingT
// интерфейсам. В качестве writer-а использующий testing.T
type Testing struct {
	testing.TB

	logger *zap.Logger

	failed   bool
	Messages []string
}

// NewTesting - logger для тестов
func NewTesting(t *testing.T, level int) slog.Logger {
	log := zaptest.NewLogger(t, zaptest.Level(CountdownLevel(level)))

	return &Testing{
		TB:     t,
		logger: log,
	}
}

// Logf - реализация logger-a testing.TB
func (z *Testing) Logf(s string, i ...interface{}) {
	z.Infof(s, i...)
}

// Fail фэйлит тест
func (z *Testing) Fail() {
	z.failed = true
}

// Failed возвращает признак того что тест завершен неуспешно
func (z *Testing) Failed() bool {
	return z.failed
}

// FailNow фэйлит тест
func (z *Testing) FailNow() {
	z.Fail()
	z.TB.FailNow()
}

// Name возвращает имя теста
func (z *Testing) Name() string {
	return z.TB.Name()
}

// Info логирует информационное сообщение
func (z *Testing) Info(args ...interface{}) {
	msg := getMessage(args...)
	z.logger.Info(msg.message, msg.fields...)
}

// Infof логирует информационное сообщение с форматированием
func (z *Testing) Infof(format string, args ...interface{}) {
	msg := getMessagef(format, args...)
	z.logger.Info(msg.message, msg.fields...)
}

// Error логирует ошибку
func (z *Testing) Error(args ...interface{}) {
	msg := getMessage(args...)
	z.logger.Error(msg.message, msg.fields...)
}

// Errorf логирует ошибку с форматированным сообщением
func (z *Testing) Errorf(format string, args ...interface{}) {
	msg := getMessagef(format, args...)
	z.logger.Error(msg.message, msg.fields...)
}

// Fatal логирует ошибку и прерывает работу
func (z *Testing) Fatal(args ...interface{}) {
	msg := getMessage(args...)
	z.logger.Fatal(msg.message, msg.fields...)
}

// Fatalf логирует ошибку и прерывает работу с форматированным сообщением
func (z *Testing) Fatalf(format string, args ...interface{}) {
	msg := getMessagef(format, args...)
	z.logger.Fatal(msg.message, msg.fields...)
}

// Warn логирует предупреждение
func (z *Testing) Warn(args ...interface{}) {
	msg := getMessage(args...)
	z.logger.Warn(msg.message, msg.fields...)
}

// Warnf логирует предупреждение с форматированным сообщением
func (z *Testing) Warnf(format string, args ...interface{}) {
	msg := getMessagef(format, args...)
	z.logger.Warn(msg.message, msg.fields...)
}

// Panic - логгирует панику
func (z *Testing) Panic(args ...interface{}) {
	msg := getMessage(args...)
	z.logger.Panic(msg.message, msg.fields...)
}

// Panicf - логгирует панику с форматированным сообщением
func (z *Testing) Panicf(format string, args ...interface{}) {
	msg := getMessagef(format, args...)
	z.logger.Panic(msg.message, msg.fields...)
}

// With - спавнит логгер с анотацией полями
func (z *Testing) With(args ...interface{}) slog.Logger {
	msg := getMessage(args...)

	newLog := &Testing{
		TB:     z.TB,
		logger: z.logger.With(msg.fields...),
	}

	return newLog
}

// Sync - синхронизация логгера
func (z *Testing) Sync() {
	_ = z.logger.Sync()
}
